```sh
time go run .
# 18446744073709551557

# real    0m1.772s
# user    0m13.271s
# sys     0m0.069s


time go run -race .
# 18446744073709551557

# real    0m14.799s
# user    1m40.147s
# sys     0m0.333s

# Escape analysis
go build -gcflags=-m=3 ./...

```