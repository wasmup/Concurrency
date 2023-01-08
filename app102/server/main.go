package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var storage atomic.Int64

func home(w http.ResponseWriter, r *http.Request) {
	n := int64(runtime.NumGoroutine())
	o := storage.Load()
	max := o
	if n > o {
		max = n
		for storage.CompareAndSwap(o, n) {
		}
	}
	log.Println(n, o, max)
	w.Write(time.Now().AppendFormat(nil, time.RFC3339))
}

func main() {
	adr := ":8080"
	if len(os.Args) > 1 {
		adr = os.Args[1]
	}

	server := http.NewServeMux()
	server.HandleFunc("/", home)

	log.Println("Starting server on", adr)
	err := http.ListenAndServe(adr, server)
	log.Fatal(err)
}
