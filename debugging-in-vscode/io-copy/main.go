package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)

	// set a breakpoint here:
	n, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
