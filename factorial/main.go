package main

import (
	"fmt"
	"math/big"
	"math/bits"
	"time"
)

func main() {
	fmt.Println(bits.Len64(factoial(22))) // 64

	t0 := time.Now()
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
