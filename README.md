# The Go Programming Language Concurrency

# Books on the Go Programming Language Concurrency 
- Concurrency in Go by Katherine Cox-Buday
- Mastering Concurrency in Go by Kozyra Nathan
- Effective Concurrency in Go by Burak Serdar


---

# Go Programming Language Books
- The Go Programming Language by Alan A. A. Donovan, and Brian W. Kernighan
- Learn Go with Tests by Chris James
- Pro Go by Adam Freeman
- Let's Go Further by Alex Edwards
- Hands-On Dependency Injection in Go by Corey Scott
- 100 Go Mistakes and How to Avoid Them by Teiva Harsanyi
- Know Go Generics by John Arundel
- Generic Data Structures and Algorithms in Go by Richard Wiener
- Go: Design Patterns for Real-World Projects by Vladimir Vivien, Mario Castro Contreras, and Mat Ryer

---

# The Go Programming Language Specification
- https://go.dev/ref/spec

# Documentation 
- https://go.dev/doc/

# Install
- https://go.dev/dl/

# The Go Playground
- https://go.dev/play/

# Tour of the Go programming language.
- https://go.dev/tour/


# Stack Overflow
- https://stackoverflow.com/tags/go/info
- https://stackoverflow.com/questions/tagged/go

# Effective Go
- https://go.dev/doc/effective_go

# Standard library
- https://pkg.go.dev/std

# Blog
- https://go.dev/blog/

# Profile-guided optimization
- https://go.dev/doc/pgo



# Memory Model
- https://go.dev/ref/mem

> "Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access."

> "To serialize access, protect the data with channel operations or other synchronization primitives such as those in the sync and sync/atomic packages."

Based on the Go memory model and [atomic package](https://pkg.go.dev/sync/atomic):

- Atomicity: Operations on atomic variables are indivisible. This means that even if multiple goroutines (lightweight threads in Go) attempt to access and modify the variable concurrently, the operation will be completed as a single unit.


- Visibility ("synchronization"): Changes made to an atomic variable by one goroutine become visible to all other goroutines immediately.


- Ordering: The atomic package provides various functions with different memory ordering options. These options control the order in which reads and writes to atomic variables are guaranteed to be seen by other goroutines. The happens-before relationship defines a partial ordering between memory accesses, ensuring that operations that "happen before" another operation are always observed in that order by other goroutines.

The documentation implies that [the package](https://pkg.go.dev/sync/atomic) provides the necessary mechanisms to ensure thread safety and atomicity.

```go
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var n uint64 = math.MaxUint64
	for ; !isPrime(n); n -= 2 {
	}
	fmt.Println(n) // 18446744073709551557
}

var nCPU = runtime.NumCPU()

func isPrime(n uint64) bool {
	if n <= 3 {
		return n == 2 || n == 3
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	var divisible atomic.Int32
	f := func(from, to uint64) {
		if divisible.Load() == 1 {
			return
		}
		if (from-5)%6 != 0 {
			from += 6 - ((from - 5) % 6) // Next multiple of 6:  5 + 6*k
		}
		for i := from; i <= to && divisible.Load() == 0; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				divisible.Store(1)
				return
			}
		}
	}

	runTasksInParallel(5, uint64(math.Sqrt(float64(n))), uint64(nCPU), f)

	return divisible.Load() == 0
}

func runTasksInParallel(from, to, n uint64, task func(uint64, uint64)) {
	totalTasks := to - from + 1
	tasksPerCore := totalTasks / n
	remainingTasks := totalTasks % n
	if tasksPerCore == 0 {
		n = remainingTasks
	}

	var wg sync.WaitGroup
	for ; n > 0; n-- {
		end := from + tasksPerCore - 1
		if remainingTasks > 0 {
			remainingTasks--
			end++
		}

		wg.Add(1)
		go func(from, to uint64) {
			defer wg.Done()
			task(from, to)
		}(from, end)

		from = end + 1
	}
	wg.Wait()
}

```

---

# HTTP Servers

## net/http 170k Requests/sec
- https://github.com/wasmup/Concurrency/tree/main/max-rps-server
- https://github.com/wasmup/Concurrency/blob/main/max-rps-server/main.go

## fasthttp 300k Requests/sec
- https://github.com/wasmup/Concurrency/tree/main/max-rps-server-fasthttp
- https://github.com/wasmup/Concurrency/blob/main/max-rps-server-fasthttp/main.go

## gnet 380k Requests/sec
- https://github.com/wasmup/Concurrency/tree/main/max-rps-server-gnet
- https://github.com/wasmup/Concurrency/blob/main/max-rps-server-gnet/main.go

---
