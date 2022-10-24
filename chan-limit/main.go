package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	limit := make(chan struct{}, 3)
	data := make(chan int)
	out := make(chan int)

	go generate(data)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calculate(data, out, limit, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	total := 0
	go func() {
		for v := range out {
			total += v
			show <- fmt.Sprint(v, " ")
		}
		close(show)
	}()

	for s := range show {
		fmt.Print(s)
		os.Stdout.Sync()
	}

	fmt.Println(" 29*30/2 =", total)
}

func generate(ch chan int) {
	for i := 0; i < 30; i++ {
		ch <- i
		time.Sleep(33 * time.Millisecond)
	}
}

func calculate(in, out chan int, limit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	limit <- struct{}{}
	show <- "{ "
	sum := <-in
	sum += <-in
	sum += <-in

	out <- sum

	<-limit
	show <- "} "
}

var show = make(chan string, 1)
