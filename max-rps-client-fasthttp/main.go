package main

import (
	"flag"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	var workers int
	var addr string
	var d time.Duration
	flag.IntVar(&workers, "c", 100, "Number of workers")
	flag.DurationVar(&d, "d", 10*time.Second, "Total duration of requests")
	flag.StringVar(&addr, "url", "http://localhost:8181", "URL")

	flag.Parse()
	fmt.Println("c", workers, "d", d)

	var ok, failed atomic.Uint64
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var c fasthttp.Client
			for time.Since(start) < d {
				makeRequest(&c, &ok, &failed, addr)
			}
		}()
	}

	wg.Wait()
	printResults(&ok, &failed, start)
}

func makeRequest(c *fasthttp.Client, ok, failed *atomic.Uint64, addr string) {
	statusCode, _, err := c.Get(nil, addr)
	if err != nil || statusCode != fasthttp.StatusOK {
		slog.Error("GET", "err", err)
		failed.Add(1)
		return
	}
	ok.Add(1)
}

func printResults(ok, failed *atomic.Uint64, start time.Time) {
	t := ok.Load()
	fmt.Println("ok", t, "failed", failed.Load())
	elapsed := time.Since(start)
	fmt.Printf("Test completed in %s\n", elapsed)
	rps := int(float64(t) / elapsed.Seconds())
	fmt.Println(rps, "Request/second")
}
