package main

import (
	"bufio"
	"encoding/binary"
	"log/slog"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

const (
	profiling = true

	n = 100_000_000

	MB         = 1024 * 1024
	bufferSize = 64 * MB
	primeLen   = 8 * MB
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))
	if profiling {
		f, err := os.Create("cpu.out")
		if err != nil {
			slog.Error("os.Create", "failed", err)
			return
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	file, err := os.Create("primes.bin")
	if err != nil {
		slog.Error("creating file", "failed", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriterSize(file, bufferSize)

	var mu sync.Mutex
	var wg sync.WaitGroup
	numCPU := runtime.NumCPU()

	chunkSize := (n + uint64(numCPU) - 1) / uint64(numCPU)
	for i := uint64(0); i < uint64(numCPU); i++ {
		start := i * chunkSize
		end := start + chunkSize - 1
		if end > n {
			end = n
		}
		wg.Add(1)
		go generatePrimes(start+1, end, writer, &mu, &wg)
	}

	wg.Wait()
	writer.Flush()

	slog.Info("Prime numbers generated and saved to primes.bin")
}

func write(primes []uint64, writer *bufio.Writer, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	b := make([]byte, 8)
	for _, prime := range primes {
		binary.LittleEndian.PutUint64(b, prime)
		_, err := writer.Write(b)
		if err != nil {
			slog.Error("writing to file", "failed", err)
			return
		}
	}
}

func generatePrimes(start, end uint64, writer *bufio.Writer, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	primes := make([]uint64, 0, primeLen)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			if len(primes) == cap(primes) {
				write(primes, writer, mu)
				primes = primes[:0]
			}
		}
	}
	if len(primes) > 0 {
		write(primes, writer, mu)
	}
}

func isPrime(n uint64) bool {
	if n <= 3 {
		return n == 3 || n == 2
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Optimized loop with 6k ± 1 optimization and early termination
	i := uint64(5)
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

// func isPrime(num uint64) bool {
// 	if num <= 3 {
// 		return num > 1
// 	}
// 	if num%2 == 0 || num%3 == 0 {
// 		return false
// 	}
// 	for i := uint64(5); i*i <= num; i += 6 {
// 		if num%i == 0 || num%(i+2) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
