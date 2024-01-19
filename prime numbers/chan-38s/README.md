```sh
	time go run .

# real    0m38.244s
# user    2m15.400s
# sys     0m4.533s

	go tool pprof -http=":8787" cpu.out

```
