# NumberToStringFast 40% Faster and no malloc

```go
// BenchmarkFormatInt-4      100000000               33.02 ns/op            7 B/op          0 allocs/op
// BenchmarkItoa-4           100000000               32.84 ns/op            7 B/op          0 allocs/op
// NumberToStringFast-4      100000000               19.47 ns/op            0 B/op          0 allocs/op
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


```sh
go test -benchtime=100000000x -benchmem -run=. -bench . my

BenchmarkFormatInt-4            100000000               33.02 ns/op            7 B/op          0 allocs/op
BenchmarkItoa-4                 100000000               32.84 ns/op            7 B/op          0 allocs/op
BenchmarkNumberToString-4       100000000               79.01 ns/op           15 B/op          1 allocs/op
BenchmarkNumberToString2-4      100000000               63.49 ns/op            7 B/op          1 allocs/op
BenchmarkNumberToString3-4      100000000               19.47 ns/op            0 B/op          0 allocs/op
```
