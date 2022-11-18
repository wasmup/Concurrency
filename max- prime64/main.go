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
	fmt.Println(nCPU)

	t0 := time.Now()
	fmt.Println(isOddPrime(18446744073709551557))
	fmt.Println(time.Since(t0)) // 7.7s
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
