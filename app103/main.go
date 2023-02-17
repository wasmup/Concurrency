package main

import "fmt"

// Given an array of integers, find two numbers such that they add up to a specific target number.
// not sorted input

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(twoSum(a, 9))
}

// O(n)
func twoSum(a []int, sum int) *two {
	m := map[int]bool{}
	for _, v := range a {
		if m[v] {
			return &two{sum - v, v}
		}
		m[sum-v] = true
	}
	return nil
}

type two struct {
	a, b int
}
