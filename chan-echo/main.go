package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 0)

	wg.Add(1)
	go echo(ch, &wg)

	log.Println("1 main tx")
	ch <- "Hi"

	log.Println("4 main rx", <-ch)

	log.Println("5 main closed")
	close(ch)

	wg.Wait()
	log.Println("7 main done")
}

func echo(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		s, ok := <-ch
		if !ok {
			log.Println("6 echo done")
			return
		}
		log.Println("2 echo", s)

		log.Println("3 echo tx")
		ch <- s
	}
}
