package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := mapTo(strconv.Itoa, a)
	fmt.Println(s)
}

func mapTo[From, To any](f func(From) To, a []From) (result []To) {
	for i := range a {
		result = append(result, f(a[i]))
	}
	return
}
