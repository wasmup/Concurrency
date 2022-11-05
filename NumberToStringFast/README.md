# NumberToStringFast 40% Faster and no malloc

```sh
go test -benchtime=100000000x -benchmem -run=. -bench . my

BenchmarkNumberToStringSt-4             100000000              287 ns/op           46 B/op          7 allocs/op
BenchmarkNumberToStringLog10-4          100000000               80 ns/op           15 B/op          2 allocs/op
BenchmarkNumberToStringLog10Unsafe-4    100000000               62 ns/op            7 B/op          1 allocs/op
BenchmarkFormatInt-4                    100000000               32 ns/op            7 B/op          0 allocs/op
BenchmarkItoa-4                         100000000               36 ns/op            7 B/op          0 allocs/op
BenchmarkNumberToStringSlice-4          100000000               19 ns/op            0 B/op          0 allocs/op
BenchmarkNumberToStringArray-4          100000000               19 ns/op            0 B/op          0 allocs/op
```

```go
// 40% faster than `strconv.FormatInt` and `strconv.Itoa`
func NumberToStringFast(n int) string {
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
```
