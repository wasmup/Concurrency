```sh
	time go run .

# real    0m21.399s
# user    1m59.538s
# sys     0m1.178s

	go tool pprof -http=":8787" cpu.out

```

<img src="21.png">