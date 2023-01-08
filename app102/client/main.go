package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	url := "http://localhost:8080"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	const max = 100
	ch := make(chan time.Duration, max)

	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t0 := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			var sb strings.Builder
			io.Copy(&sb, resp.Body)

			t1, err := time.Parse(time.RFC3339, sb.String())
			if err != nil {
				log.Fatal(err)
			}
			ch <- t0.Sub(t1)
		}()
	}

	wg.Wait()
	close(ch)
	for t := range ch {
		fmt.Println(t)
	}
}
