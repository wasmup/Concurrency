package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var t0 = time.Now()
	fmt.Println(isOddPrime(18446744073709551557)) // 2.784962976s
	fmt.Println(time.Since(t0))

	t0 = time.Now()
	fmt.Println(isBigOddPrime(18446744073709551557)) // 8.633154118s
	fmt.Println(time.Since(t0))

	t0 = time.Now()
	fmt.Println(isBigOddPrime2(18446744073709551557)) // 13.985002732s
	fmt.Println(time.Since(t0))
}

var nCPU = runtime.NumCPU()

func isOddPrime(n uint64) bool {
	q := uint64(math.Sqrt(float64(n)))
	step := q / uint64(nCPU)
	if step&1 == 1 {
		step++ // make it even
	}
	var wg sync.WaitGroup
	var quit int32
	start := uint64(3)
	for k := 0; k < nCPU; k++ {
		end := start + step
		if end > q {
			end = q
		}
		wg.Add(1)
		go func(start, end uint64) {
			defer wg.Done()
			for i := start; i <= end && atomic.LoadInt32(&quit) == 0; i += 2 {
				if n%i == 0 {
					atomic.StoreInt32(&quit, 1)
					return
				}
			}
		}(start, end)
		start = end + 2
	}
	wg.Wait()
	return quit == 0
}

func isBigOddPrime(n uint64) bool {
	if n%3 == 0 {
		return false
	}
	q := uint64(math.Sqrt(float64(n)))
	for i := uint64(5); i <= q; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func isBigOddPrime2(n uint64) bool {
	if n%3 == 0 {
		return false
	}
	q := uint64(math.Sqrt(float64(n)))
	for i := uint64(5); i <= q; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
