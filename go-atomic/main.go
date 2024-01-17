package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	const n = 32
	const mx = uint64(1 << 32)
	fmt.Println("max =", mx) // max = 4294967296

	numCPUs := runtime.NumCPU()
	chunkSize := mx / uint64(numCPUs)

	// file, err := os.Create("output.bin")
	// if err != nil {
	// 	fmt.Println("Error creating output file:", err)
	// 	return
	// }
	// defer file.Close()

	for i := 0; i < numCPUs; i++ {
		start := uint64(i) * chunkSize
		end := start + chunkSize
		if i == numCPUs-1 {
			end = mx
		}

		wg.Add(1)
		go processChunk(start, end, n, nil)
	}

	wg.Wait()
	fmt.Println(total.Load())
}

var (
	// mu    sync.Mutex
	total atomic.Uint64
	wg    sync.WaitGroup
)

func processChunk(start, end uint64, n int, file *os.File) {
	defer wg.Done()
	for i := start; i < end; i++ {
		if fn(i) {
			total.Add(1)
			// mu.Lock()
			// err := binary.Write(file, binary.LittleEndian, i)
			// mu.Unlock()
			// if err != nil {
			// 	fmt.Println("Error writing to file:", err)
			// 	return
			// }
		}
	}
}

func fn(u uint64) bool {
	n0 := (u & 0x1) != 0
	n1 := (u & 0x2) != 0
	n2 := (u & 0x4) != 0
	n3 := (u & 0x8) != 0
	n4 := (u & 0x10) != 0
	n5 := (u & 0x20) != 0
	n6 := (u & 0x40) != 0
	n7 := (u & 0x80) != 0
	n8 := (u & 0x100) != 0
	n9 := (u & 0x200) != 0
	n10 := (u & 0x400) != 0
	n11 := (u & 0x800) != 0
	n12 := (u & 0x1000) != 0
	n13 := (u & 0x2000) != 0
	n14 := (u & 0x4000) != 0
	n15 := (u & 0x8000) != 0
	n16 := (u & 0x10000) != 0
	n17 := (u & 0x20000) != 0
	n18 := (u & 0x40000) != 0
	n19 := (u & 0x80000) != 0
	n20 := (u & 0x100000) != 0
	n21 := (u & 0x200000) != 0
	n22 := (u & 0x400000) != 0
	n23 := (u & 0x800000) != 0
	n24 := (u & 0x1000000) != 0
	n25 := (u & 0x2000000) != 0
	n26 := (u & 0x4000000) != 0
	n27 := (u & 0x8000000) != 0
	n28 := (u & 0x10000000) != 0
	n29 := (u & 0x20000000) != 0
	n30 := (u & 0x40000000) != 0
	n31 := (u & 0x80000000) != 0

	return (((((n0 && n1) && (n2 || n3)) || ((n4 && n5) && (n6 && n7))) && (((n8 || n9) || (n10 && n11)) && ((n12 || n13) || (n14 || n15)))) || ((((n16 || n17) || (n18 || n19)) && ((n20 || n21) && (n22 && n23))) || (((n24 && n25) && (n26 && n27)) && ((n28 && n29) || (n30 && n31)))))
}
