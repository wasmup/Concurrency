package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var t0 = time.Now()
	log.Println(1, fib(1))
	log.Println(2, fib(2))
	log.Println(3, fib(3))

	a, b := 41, 42
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Println(a, fib(a), time.Since(t0))
		// 41 102334155 472.232595ms
		wg.Done()
	}()
	log.Println(b, fib(b), time.Since(t0))
	// 42 165580141 763.482716ms

	wg.Wait()
	log.Println(time.Since(t0)) // 763.497185ms
}

func fib(n int) int {
	switch n {
	case 1, 2:
		return n - 1
	}
	return fib(n-1) + fib(n-2)
}
