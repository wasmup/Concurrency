package fmt

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var m sync.Mutex
var dLock, dTotal int64

func Print(s string) {
	t0 := time.Now()

	if v := atomic.AddInt32(&token, +1); v > atomic.LoadInt32(&max) {
		atomic.StoreInt32(&max, v)
	}

	m.Lock()
	t1 := time.Now()
	fmt.Print(s)
	m.Unlock()
	atomic.AddInt64(&dLock, int64(time.Since(t1)))

	atomic.AddInt32(&token, -1)
	atomic.AddInt64(&dTotal, int64(time.Since(t0)))
}

var token, max int32

func ShowMax() {
	fmt.Println("\n max =", atomic.LoadInt32(&max))
	fmt.Println("\n dLock =", time.Duration(atomic.LoadInt64(&dLock)))
	fmt.Println("\n dTotal =", time.Duration(atomic.LoadInt64(&dTotal)))
}
