package main

import (
	"strconv"
	"testing"
)

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(int64(i), 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkNumberToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToString(i)
	}
}

func BenchmarkNumberToString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToString2(i)
	}
}

func BenchmarkNumberToStringFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToStringFast(i)
	}
}
