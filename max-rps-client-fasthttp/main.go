package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	var workers, requestsPerWorker int
	flag.IntVar(&workers, "c", 32, "Number of workers")
	flag.IntVar(&requestsPerWorker, "n", 100000, "Total number of requests")
	flag.Parse()
	requestsPerWorker /= workers
	fmt.Println("c", workers, "rpw", requestsPerWorker, "n", workers*requestsPerWorker)

	var ok, failed atomic.Uint64
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < requestsPerWorker; i++ {
				makeRequest(&ok, &failed)
			}
		}()
	}

	wg.Wait()
	printResults(&ok, &failed, start)
}

func makeRequest(ok, failed *atomic.Uint64) {
	statusCode, _, err := fasthttp.Get(nil, "http://localhost:8080/")
	if err != nil || statusCode != fasthttp.StatusOK {
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
