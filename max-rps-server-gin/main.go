package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	requestCount int
	elapsedTime  time.Duration
	mu           sync.Mutex
)

func main() {
	gin.SetMode(gin.ReleaseMode) // Set Gin to release mode

	go printStats()

	router := gin.New()
	router.Use(gin.Recovery()) // Disable detailed error information

	router.GET("/", handler)

	// Listen and serve on 0.0.0.0:8080
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func handler(c *gin.Context) {
	t0 := time.Now()

	q := c.Query("q")
	if q == "" {
		q = "Hi"
	}
	c.String(http.StatusOK, base64.StdEncoding.EncodeToString([]byte(q)))

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
