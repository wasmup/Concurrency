package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	go printStats()

	http.HandleFunc("/", handler)
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	q := r.URL.Query().Get("q")
	if q == "" {
		q = "Hi"
	}
	w.Write([]byte(base64.StdEncoding.EncodeToString([]byte(q))))

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
