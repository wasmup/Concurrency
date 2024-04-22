package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mtx sync.Mutex
	cond := sync.NewCond(&mtx)
	started := false
	const m = 10
	d := make([]Worker, m)
	var wg sync.WaitGroup
	var t0 = time.Now()
	var t1 time.Time

	for i := 0; i < m; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			d[n].start = time.Since(t0)
			mtx.Lock()
			for !started {
				cond.Wait()
			}
			mtx.Unlock()
			d[n].stop = time.Since(t1)

		}(i)
	}

	mtx.Lock()
	started = true
	t1 = time.Now()
	cond.Broadcast() // Wake up all waiting goroutines
	mtx.Unlock()

	wg.Wait()

	for i, w := range d {
		fmt.Println("Worker", i, "duration since Wake up call Broadcast:", w.stop, "duration since t0:", w.start)
	}
}

type Worker struct {
	start time.Duration
	stop  time.Duration
}
