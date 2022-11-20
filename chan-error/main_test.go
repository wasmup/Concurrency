package main

import "testing"

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
