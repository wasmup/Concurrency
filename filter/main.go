package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := filter(odd, a)
	fmt.Println(result)
}

func filter[T any](f func(T) bool, a []T) (result []T) {
	for i := range a {
		if f(a[i]) {
			result = append(result, a[i])
		}
	}
	return
}

func odd(i int) bool {
	return i&1 == 1
}
func even(i int) bool {
	return i&1 == 0
}
