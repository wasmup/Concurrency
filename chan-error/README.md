# error retrn is 23% faster than chan error

go test -benchtime=10000000x -benchmem -run=. -bench . my

BenchmarkErrorReturn-4          10000000               166.7 ns/op            40 B/op          2 allocs/op
BenchmarkErrorChan-4            10000000               217.9 ns/op            40 B/op          2 allocs/op

ok      my      3.851s