```sh
	time go run .

# real    0m20.627s
# user    1m53.853s
# sys     0m0.240s

	go tool pprof -http=":8787" cpu.out

```

<img src="21.png">