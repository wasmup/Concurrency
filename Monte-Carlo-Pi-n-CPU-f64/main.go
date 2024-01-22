package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const total = 1e8

var (
	count int
	mu    sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	n := runtime.NumCPU()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go monteCarloPi(i, n, &wg)
	}

	wg.Wait()

	pi := float64(count) / float64(total) * 4
	fmt.Println("Approximated Pi:", pi)
}

func monteCarloPi(seed, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	localCount := 0
	rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(seed)))

	for i := 0; i < total/n; i++ {
		x := rng.Float64()
		y := rng.Float64()

		if x*x+y*y < 1.0 {
			localCount++
		}
	}

	mu.Lock()
	count += localCount
	mu.Unlock()
}
