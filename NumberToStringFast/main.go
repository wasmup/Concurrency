package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	fmt.Println(NumberToStringSt(-1000000))

	fmt.Println(len(fmt.Sprint(math.MaxInt64))) // 19
	fmt.Println(len(fmt.Sprint(math.MinInt64))) // 20

	fmt.Println(NumberToStringArray(-1000000))
	fmt.Println(NumberToStringArray(-42))

	const N = 100000000
	t0 := time.Now()
	for i := 0; i < N; i++ {
		_ = strconv.FormatInt(int64(i), 10)
	}
	fmt.Println(time.Since(t0)) // 3.2

	t0 = time.Now()
	for i := 0; i < N; i++ {
		_ = strconv.Itoa(i)
	}
	fmt.Println(time.Since(t0)) // 3.2s

	t0 = time.Now()
	for i := 0; i < N; i++ {
		_ = NumberToStringArray(i)
	}
	fmt.Println(time.Since(t0)) // 1.85s
}

func NumberToStringLog10(n int) string {
	if n == 0 {
		return "0"
	}
	m := 0
	if n < 0 {
		n = -n
		m++
	}
	m += (int)(math.Log10(float64(n)) + 1)
	b := make([]byte, m)
	i := m - 1
	for ; n != 0; n /= 10 {
		b[i] = byte(n%10) + '0'
		i--
	}
	if i == 0 {
		b[0] = '-'
	}
	return string(b)
}

func NumberToStringLog10Unsafe(n int) string {
	if n == 0 {
		return "0"
	}
	m := 0
	if n < 0 {
		n = -n
		m++
	}
	m += (int)(math.Log10(float64(n)) + 1)
	b := make([]byte, m)
	i := m - 1
	for ; n != 0; n /= 10 {
		b[i] = byte(n%10) + '0'
		i--
	}
	if i == 0 {
		b[0] = '-'
	}
	return *(*string)(unsafe.Pointer(&b))
}

func NumberToStringArray(n int) string {
	if n == 0 {
		return "0"
	}
	sign := n < 0
	if sign {
		n = -n
	}
	var b [20]byte
	i := 20
	for ; n != 0; n /= 10 {
		i--
		b[i] = byte(n%10) + '0'
	}
	if sign {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

func NumberToStringSlice(n int) string {
	if n == 0 {
		return "0"
	}
	sign := n < 0
	if sign {
		n = -n
	}
	b := make([]byte, 20)
	i := len(b)
	for ; n != 0; n /= 10 {
		i--
		b[i] = byte(n%10) + '0'
	}
	if sign {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

func NumberToStringSt(n int) (s string) {
	if n == 0 {
		return "0"
	}
	sign := n < 0
	if sign {
		n = -n
	}
	for ; n != 0; n /= 10 {
		s = string(byte(n%10)+'0') + s
	}
	if sign {
		s = "-" + s
	}
	return
}
