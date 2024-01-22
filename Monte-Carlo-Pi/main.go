package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	const total = 1e8
	x := make([]float64, total)
	count := 0
	for i := range x {
		if search(rand.Intn(r), rand.Intn(r)) {
			count++
		}
		x[i] = 4 * float64(count) / float64(i+1)
	}

	u := mean(x)
	sig := standardDeviation(x, u)
	fmt.Println("mean =", u, "sigma =", sig, "standard error =", standardError(sig, len(x)))

	fmt.Println("Pi =", math.Pi, "Pi-mean =", math.Pi-u)

}

func standardError(sig float64, n int) float64 {
	return sig / math.Sqrt(float64(n))
}

func standardDeviation(xs []float64, ave float64) float64 {
	var d, sum float64
	for _, x := range xs {
		d = x - ave
		sum += d * d
	}
	return math.Sqrt(sum / float64(len(xs)-1))
}

func mean(xs []float64) (ave float64) {
	for _, x := range xs {
		ave += x
	}
	ave /= float64(len(xs))
	return
}

const r = 1e9 // math.MaxInt = 9223372036854775807

func search(x, y int) bool {
	return x*x+y*y < r*r
}
