package main

import (
	"fmt"
	"os"
)

func main() {
	max := 4
	if len(os.Args) > 1 {
		fmt.Sscan(os.Args[1], &max)
	}
	for _, v := range All(max) {
		fmt.Println(v)
	}
}

type queens struct {
	m   int
	all []int
}

func All(n int) [][]int {
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1
	}
	p := &queens{n, a}
	return p.queens(n)
}

func (p *queens) queens(n int) (result [][]int) {
	if n == 1 {
		result = make([][]int, p.m)
		for i := range result {
			result[i] = []int{i + 1}
		}
		return
	}
	for _, columns := range p.queens(n - 1) {
		for _, row := range diff(p.all, columns) {
			if ok(row, 1, columns) {
				result = append(result, append([]int{row}, columns...))
			}
		}
	}
	return
}

func ok(row, n int, columns []int) bool {
	if len(columns) == 0 {
		return true
	}
	return row != columns[0]+n && row != columns[0]-n && ok(row, n+1, columns[1:])
}

func diff(a, b []int) (result []int) {
	m := make(map[int]bool, len(b))
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}
