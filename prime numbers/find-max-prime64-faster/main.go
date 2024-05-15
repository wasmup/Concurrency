package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Sscan(os.Args[1], &nCPU)
	}

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
	f := func(from, to uint64) {
		if divisible.Load() == 1 {
			return
		}
		if (from-5)%6 != 0 {
			from += 6 - ((from - 5) % 6) // Next multiple of 6:  5 + 6*k
		}
		for i := from; i <= to && divisible.Load() == 0; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				divisible.Store(1)
				return
			}
		}
	}

	runTasksInParallel(5, uint64(math.Sqrt(float64(n))), uint64(nCPU), f)

	return divisible.Load() == 0
}

// from >= to
// n >= 1
func runTasksInParallel(from, to, n uint64, task func(uint64, uint64)) {
	if n == 1 || from == to {
		task(from, to)
		return
	}

	totalTasks := to - from + 1
	tasksPerCore := totalTasks / n
	remainingTasks := totalTasks % n
	if tasksPerCore == 0 {
		n = remainingTasks
	}

	var wg sync.WaitGroup
	for ; n > 0; n-- {
		end := from + tasksPerCore - 1
		if remainingTasks > 0 {
			remainingTasks--
			end++
		}

		wg.Add(1)
		go func(from, to uint64) {
			defer wg.Done()
			task(from, to)
		}(from, end)

		from = end + 1
	}
	wg.Wait()
}
