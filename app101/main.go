package main

import (
	"app101/fmt"
	"math/rand"
	"sync"
	"time"
)

// var fmt = log.New(os.Stdout, "app101/", log.Lshortfile) // thread safe

// var fmt = log.New(os.Stdout, "", 0) // thread safe

const max = 100

func main() {
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	wg.Add(1)
	go longRunningTask(&wg)

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func() {
			fmt.Print(".")
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.ShowMax()
	// max = 123
	// dLock = 3.331535ms
	// dTotal = 485.141125ms
}

func longRunningTask(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func() {
			fmt.Print("_")
			wg.Done()
		}()
	}
}
