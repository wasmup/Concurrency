package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890
// -t4: Specifies the number of threads to use for concurrent requests. In this case, it will create 4 threads, each capable of making requests simultaneously.
// -c80: Sets the concurrency level, meaning the number of simultaneous connections to maintain. Here, wrk will try to keep 80 connections open at once.
// -d10s: Determines the duration of the test. The server will be benchmarked for 10 seconds (10s).

func main() {
	var workers, requestsPerWorker int
	flag.IntVar(&workers, "c", 32, "Number of workers")
	flag.IntVar(&requestsPerWorker, "n", 100000, "Total number of requests")
	flag.Parse()
	requestsPerWorker /= workers
	fmt.Println("c", workers, "rpw", requestsPerWorker, "n", workers*requestsPerWorker)

	// f, err := os.Create("cpu.pprof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	var ok, failed atomic.Uint64
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
			if err != nil {
				log.Fatal(err)
			}
			var client = &http.Client{
				Transport: &http.Transport{
					MaxIdleConnsPerHost: requestsPerWorker,
				},
			}
			for i := 0; i < requestsPerWorker; i++ {
				makeRequest(client, req, &ok, &failed)
			}
		}()
	}

	wg.Wait()
	printResults(&ok, &failed, start)
}

func makeRequest(client *http.Client, req *http.Request, ok, failed *atomic.Uint64) {
	resp, err := client.Do(req)
	if err != nil {
		failed.Add(1)
		return
	}
	resp.Body.Close()
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
