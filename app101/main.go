package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func longRunningTaskWithSomeResult(ctx context.Context, wg *sync.WaitGroup, ch chan Result) {
	defer wg.Done()

	for i := 0; i < 10 && ctx.Err() == nil; i++ { // termination condition.

		time.Sleep(200 * time.Millisecond) // long running task

		fmt.Print(".") // race
	}

	log.Println("long running task is done")
	ch <- Result{42, ctx.Err()} // blocking ?   g:leak ?
}

func main() {
	var wg sync.WaitGroup
	var ch = make(chan Result, 2) // edit

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	wg.Add(1)
	go longRunningTaskWithSomeResult(ctx, &wg, ch)

	wg.Add(1)
	go longRunningTaskWithSomeResult(ctx, &wg, ch) // what happens

	go func() { // superfluous
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond) // long running task
			fmt.Print("-")                     // race
		}
	}()

	log.Println("\nwait:")
	wg.Wait()

	result := <-ch
	if result.err != nil {
		log.Println(result.err)
		return // note
	}
	log.Println(result.value)
}

type Result struct {
	value int
	err   error
}
