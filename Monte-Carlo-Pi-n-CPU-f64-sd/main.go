package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	n := runtime.NumCPU()
	counts := make([]int, n)
	const total = 1e8
	for i := 0; i < n; i++ {
		wg.Add(1)
		go monteCarloPi(counts, i, total/n, &wg)
	}

	wg.Wait()

	mean, stdDev := meanStdDev(counts)
	standardError := stdDev / math.Sqrt(float64(n))

	fmt.Println("Mean Pi:", 4*mean/total*float64(n))
	fmt.Println("Standard Deviation of Pi estimates:", 4*stdDev/total*float64(n))
	fmt.Println("Standard Error of the Mean:", 4*standardError/total*float64(n))
}

func monteCarloPi(counts []int, i, total int, wg *sync.WaitGroup) {
	defer wg.Done()

	rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
	count := 0
	for ; total > 0; total-- {
		x := rng.Float64()
		y := rng.Float64()

		if x*x+y*y < 1.0 {
			count++
		}
	}

	counts[i] = count
}

func meanStdDev(data []int) (mean, stdDev float64) {
	sum := 0
	for _, value := range data {
		sum += value
	}
	mean = float64(sum) / float64(len(data))

	var sumSquaredDiff float64
	for _, value := range data {
		diff := float64(value) - mean
		sumSquaredDiff += diff * diff
	}
	stdDev = math.Sqrt(sumSquaredDiff / float64(len(data)-1))

	return
}
