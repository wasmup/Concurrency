package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var n uint64 = math.MaxUint64
	for ; !isPrime(n); n -= 2 {
	}
	fmt.Println(n) // 18446744073709551557
}

var nCPU = runtime.NumCPU()

func isPrime(n uint64) bool {
	if n <= 3 {
		return n == 2 || n == 3
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	var divisible atomic.Int32
	q := uint64(math.Sqrt(float64(n)))
	step := q / uint64(nCPU)
	step += step & 1 // make it even
	var wg sync.WaitGroup
	var start uint64 = 3

	for k := 0; k < nCPU; k++ {
		end := start + step
		if end > q {
			end = q
		}

		wg.Add(1)
		go func(start, end uint64) {
			defer wg.Done()
			for i := start; i <= end && divisible.Load() == 0; i += 2 {
				if n%i == 0 {
					divisible.Store(1)
					return
				}
			}
		}(start, end)

		start = end + 2
	}

	wg.Wait()

	return divisible.Load() == 0
}
