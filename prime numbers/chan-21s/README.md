```sh
time go run .

# real    0m21.016s
# user    1m57.014s
# sys     0m1.056s

go tool pprof -http=":8787" cpu.out

```

<img src="21.png">