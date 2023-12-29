package main

import (
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	go printStats()

	if err := fasthttp.ListenAndServe(":8080", handler); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func handler(ctx *fasthttp.RequestCtx) {
	t0 := time.Now()

	q := string(ctx.QueryArgs().Peek("q"))
	if q == "" {
		q = "Hi"
	}
	ctx.Write([]byte(base64.StdEncoding.EncodeToString([]byte(q))))

	d := time.Since(t0)
	mu.Lock()
	requestCount++
	elapsedTime += d
	mu.Unlock()
}

func printStats() {
	for {
		time.Sleep(time.Second)
		mu.Lock()
		n := requestCount
		d := elapsedTime
		mu.Unlock()
		fmt.Printf("Requests: %d, Total Elapsed Time: %s\n", n, d)
	}
}

var (
	requestCount int
	elapsedTime  time.Duration
	mu           sync.Mutex
)
