package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 1)

	go longRunningTask(ch)

	fmt.Println("Another task:")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond) // long running task
		fmt.Print("-")
	}

	fmt.Println("\nwait:")
	result := <-ch
	fmt.Println(result)
}

func longRunningTask(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond) // long running task
		fmt.Print(".")
	}
	result := 42
	fmt.Println("long running task is done")
	ch <- result
}
