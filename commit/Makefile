Version = "1.0.0"
Name = "MyApp"
GOOS = linux
GOARCH = amd64

GIT_COMMIT := $(shell git rev-parse HEAD)

build:
	CGO_ENABLED=0 go build -ldflags="-s -X main.Version=$(Version) -X main.Name=$(Name) -X main.GitCommit=$(GIT_COMMIT)" -trimpath=true .
	go version -m ./commit

run:
	./commit

clean:
	rm -rf commit

init:
	go get -x -u all
	go mod tidy -x
	go mod vendor

info:
	ls -lh ./commit
	ls -l ./commit
	strip -s ./commit 
	ls -l ./commit
	file ./commit
	go version -m ./commit
	go version -m $(which go)

.PHONY: build run clean info

