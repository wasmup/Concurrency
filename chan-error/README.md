# error retrn is 1.25x faster than chan error
# error retrn is 20x faster than goroutine
```sh
BenchmarkErrorReturn-4                   1000000               183.0 ns/op            40 B/op          2 allocs/op
BenchmarkErrorChan-4                     1000000               230.4 ns/op            40 B/op          2 allocs/op
BenchmarkGoroutineErrorChan-4            1000000              3796 ns/op             256 B/op         11 allocs/op

ok      my      4.217s



go test -benchtime=4s -benchmem -run=. -bench . my

BenchmarkErrorReturn-4                  26622122               162.6 ns/op            40 B/op          2 allocs/op
BenchmarkErrorChan-4                    20648920               217.6 ns/op            40 B/op          2 allocs/op
BenchmarkGoroutineErrorChan-4            1353189              3496 ns/op             256 B/op         11 allocs/op

ok      my      17.569s


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
