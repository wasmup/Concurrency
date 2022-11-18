package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	var nc chan int
	ch := make(chan int, 1)
	ch <- 1

	select {
	case v, ok := <-nc:
		fmt.Println(v, ok)

	case v, ok := <-ch:
		fmt.Println(v, ok)
	}

	close(ch)

	select {
	case v, ok := <-nc:
		fmt.Println(v, ok)

	case v, ok := <-ch:
		fmt.Println(v, ok)
	}

	fmt.Println(runtime.NumGoroutine())
	go task1()

	fmt.Println(runtime.NumGoroutine())
	go task2()

	fmt.Println(runtime.NumGoroutine())
	go task3()

	fmt.Println(runtime.NumGoroutine())

	ch = make(chan int, 1) // comment ?
	ch <- 1
	v, ok := <-ch
	fmt.Println(v, ok)

	close(ch)
	fmt.Println(<-ch) // read from a closed channel is a non-blocking
	v, ok = <-ch
	fmt.Println(v, ok)

	go gen()
	go fanOut()
	go fanIn()
	for v := range d {
		fmt.Println(v)
	}
}

func task1() {
	var ch chan int
	log.Println("1")
	<-ch // read from a nil chan is blocking (0% CPU usage)
	log.Println("2")
}

func task2() {
	var ch chan int
	log.Println("1")
	ch <- 0 // write to a nil chan is blocking (0% CPU usage)
	log.Println("2")
}

func task3() {

	log.Println("3")
	select {} // (0% CPU usage)
}

var (
	a = make(chan int)
	b = make(chan int)
	c = make(chan int)
	d = make(chan int)
)

func gen() {
	for i := 1; i < 3; i++ {
		a <- i
	}
	close(a)
}
func fanOut() {
	for v := range a {
		b <- v
		c <- 10 * v
	}
	close(b)
	close(c)
}
func fanIn() {
	for b != nil || c != nil {
		select {
		case v, ok := <-b:
			if !ok {
				b = nil
				break
			}
			d <- v

		case v, ok := <-c:
			if ok {
				d <- v
			} else {
				c = nil
			}
		}
	}
	close(d)
}
