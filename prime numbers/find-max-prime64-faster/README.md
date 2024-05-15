```sh
time go run . 1
# 18446744073709551557
# real	0m9.633s
# user	0m9.665s
# sys	0m0.093s

 time go run . 
# 18446744073709551557
# real	0m1.744s
# user	0m13.232s
# sys	0m0.045s


time go run -race .
# 18446744073709551557
# real	0m19.762s
# user	1m49.617s
# sys	0m1.125s


# Escape analysis
go build -gcflags=-m=3 ./...

```