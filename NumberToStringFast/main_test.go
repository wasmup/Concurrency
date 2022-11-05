package main

import (
	"strconv"
	"testing"
)

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkNumberToStringLog10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToStringLog10(i)
	}
}

func BenchmarkNumberToStringLog10Unsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToStringLog10Unsafe(i)
	}
}

func BenchmarkNumberToStringSt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToStringSt(i)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(int64(i), 10)
	}
}

func BenchmarkNumberToStringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToStringSlice(i)
	}
}

func BenchmarkNumberToStringArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumberToStringArray(i)
	}
}
