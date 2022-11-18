package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	b, err := io.ReadAll(os.Stdin) // Ctrl+D to finish in terminal
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b)
	fmt.Println(len(b))
}
