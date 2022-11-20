# error retrn is 1.25x faster than chan error
# error retrn is 20x faster than goroutine
```sh
	go run .
1.664232336s
2.237619654s
37.535908032s
1.582753946s // fmt.Errorf("parachute failed: %d", id).Error()

1.817283284s
2.915837246s
35.890209896s
483.925435ms	// return "parachute failed: " + strconv.Itoa(id)

go test -benchtime=10000000x -benchmem -run=. -bench . my

BenchmarkStringReturn-4                 10000000               174.2 ns/op            40 B/op          2 allocs/op
BenchmarkErrorReturn-4                  10000000               173.7 ns/op            40 B/op          2 allocs/op
BenchmarkErrorChan-4                    10000000               222.2 ns/op            40 B/op          2 allocs/op
BenchmarkGoroutineErrorChan-4           10000000              3670 ns/op             256 B/op         11 allocs/op

ok      my      42.407s


# Itoa 4x faster than Errorf
BenchmarkItoa-4   	20609462	        53.02 ns/op	       7 B/op	       0 allocs/op
BenchmarkErrorf-4    6337472	       227.3 ns/op	      54 B/op	       3 allocs/op
```

```go
func f2(id int) error {
	id++
	return f1(id)
}
func f1(id int) error {
	id++
	return fmt.Errorf("parachute failed: %d", id)
}
```

```go
func e2(id int, ch chan error) {
	id++
	e1(id, ch)
}
func e1(id int, ch chan error) {
	id++
	ch <- fmt.Errorf("parachute failed: %d", id)
}

```


```go
func g2(id int, ch chan error) {
	id++
	go g1(id, ch)
}
func g1(id int, ch chan error) {
	id++
	ch <- fmt.Errorf("parachute failed: %d", id)
}
```
