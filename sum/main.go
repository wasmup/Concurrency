package main

import "fmt"

func main() {
	fmt.Println(sum(1, 100))
	fmt.Println(sum2(1, 100))
}

func sum(a, b int) int {
	return (a + b) * (b - a + 1) / 2
}

func sum2(a, b int) int {
	s := 0
	for i := a; i <= b; i++ {
		s += i
	}
	return s
}
