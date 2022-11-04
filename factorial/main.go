package main

import (
	"fmt"
	"math"
	"math/big"
	"math/bits"
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

	fmt.Println(bits.Len64(factoial(22))) // 64

	t0 = time.Now()
	f := factoialBig(100_000)
	fmt.Println(time.Since(t0)) // 1s

	fmt.Println(f.BitLen()) // 1516705
}

func factoial(n uint64) (f uint64) {
	f = 1
	for i := n; i > 1; i-- {
		f *= i
	}
	return
}

func factoialBig(n int64) (f *big.Int) {
	f = big.NewInt(1)
	for i := n; i > 1; i-- {
		f = f.Mul(f, big.NewInt(i))
	}
	return
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
