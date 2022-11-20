package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkStringReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = c10(1)
	}
}

func BenchmarkErrorReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = f10(1)
	}
}

func BenchmarkErrorChan(b *testing.B) {
	ch := make(chan error, 1)
	for i := 0; i < b.N; i++ {
		e10(1, ch)
		_ = <-ch
	}
}

func BenchmarkGoroutineErrorChan(b *testing.B) {
	ch := make(chan error, 1)
	for i := 0; i < b.N; i++ {
		g10(1, ch)
		_ = <-ch
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "parachute failed: " + strconv.Itoa(i)
	}
}
func BenchmarkErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Errorf("parachute failed: %d", i).Error()
	}
}
