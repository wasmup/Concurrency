package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("A: Started")
		time.Sleep(1 * time.Second)
		fmt.Println("A: Finished")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("B: Started")
		time.Sleep(1 * time.Second)
		fmt.Println("B: Finished")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("C: Started")
		time.Sleep(1 * time.Second)
		fmt.Println("C: Finished")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("D: Started")
		time.Sleep(1 * time.Second)
		fmt.Println("D: Finished")
	}()

	wg.Wait()
}
