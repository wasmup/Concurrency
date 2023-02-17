package main

import (
	"fmt"
)

// Given an array of integers, find two numbers such that they add up to a specific target number.
// sorted input

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(twoSum(a, 9))
}

// O(n)
func twoSum(a []int, t int) *two {
	for i, j := 0, len(a)-1; i < j; {
		s := a[i] + a[j]
		switch {
		case t == s:
			return &two{a[i], a[j]}
		case s < t:
			i++
		default:
			j--
		}
	}
	return nil
}

type two struct {
	a, b int
}
