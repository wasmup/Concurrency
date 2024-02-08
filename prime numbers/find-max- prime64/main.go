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
	fmt.Println("NumCPU:", nCPU)

	t0 := time.Now()
	var n uint64 = math.MaxUint64
	for i := 0; !isOddPrime(n); n -= 2 {
		i++
		fmt.Println(i, n)
	}
	fmt.Println("u64 max prime:", n)       // 18446744073709551557
	fmt.Println("diff:", math.MaxUint64-n) // 58
	fmt.Println(time.Since(t0))            // 3.287094056s

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
