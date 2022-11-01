package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	t0 := time.Now()

	const Timeout = 1 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	const key = "sum"

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum := 0
		for i := 1; ctx.Err() == nil; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(i, time.Since(t0))
			sum += i // example job
		}
		ctx = context.WithValue(ctx, key, sum) // beware of concurrent read while write to ctx it is not safe
	}()

	wg.Wait()
	fmt.Println(key, ctx.Value(key))
	fmt.Println("done.")
}
