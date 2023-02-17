package main

import (
	"fmt"
	"sort"
)

// Given an array of integers, find two numbers such that they add up to a specific target number.
// sorted input

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(twoSum(a, 9))
}

// O(n log n)
func twoSum(a []int, sum int) *two {
	for _, v := range a {
		i := sort.Search(len(a), func(i int) bool { return a[i] == v }) // O(log n)
		if i > 0 && i < len(a) {
			return &two{sum - v, v}
		}
	}
	return nil
}

type two struct {
	a, b int
}
