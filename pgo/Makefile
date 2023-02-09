all:
	go mod init example.com/markdown
	go get gitlab.com/golang-commonmark/markdown@bf3e522c626a
	go mod tidy -v
	go build -v -o markdown.nopgo
	ls -lh markdown.nopgo

	curl -o README.md -L "https://raw.githubusercontent.com/golang/go/c16c2c49e2fa98ae551fc6335215fadd62d33542/README.md"

	gnome-terminal -- go  run .
	curl -o README.html --data-binary @README.md http://localhost:8080/render

	go run load/main.go

# While that is running, download a profile from the server:
	curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
# Once this completes, kill the load generator and the server.
	mv cpu.pprof default.pgo
	go build -pgo=auto -o markdown.withpgo

	./markdown.nopgo
	go test ./load -bench=. -count=20 -source ../README.md > nopgo.txt
# Once that completes, kill the original server and start the version with PGO:
	./markdown.withpgo
	go test ./load -bench=. -count=20 -source ../README.md > withpgo.txt

	go install -v golang.org/x/perf/cmd/benchstat@latest
	benchstat nopgo.txt withpgo.txt
#        │  nopgo.txt  │            withpgo.txt             │
#        │   sec/op    │   sec/op     vs base               │
# Load-4   288.6µ ± 4%   263.1µ ± 5%  -8.82% (p=0.000 n=20)
# 288.6/263.1=1.096921323 => 9.6% faster


# Profile-guided optimization 
# https://go.dev/doc/pgo
