VERSION = "1.0.1"
NAME = "MyApp"
GOOS = linux
GOARCH = amd64

GIT_COMMIT := $(shell git rev-parse HEAD)

all: build

build:
	CGO_ENABLED=0 go build -trimpath=true -ldflags="-s -X main.Version=$(VERSION) -X main.Name=$(NAME) -X main.GitCommit=$(GIT_COMMIT)"
	go version -m ./app

run:
	go run -trimpath=true .

clean:
	rm -rf app

init:
	go get -x -u all
	go mod tidy -x
	go mod vendor

info:
	ls -lh ./app
	ls -l ./app
	strip -s ./app 
	ls -l ./app
	file ./app
	go version -m ./app
	go version -m $(which go)


	go tool pprof -http=":8789" cpu.out
	go tool pprof -http=":8787" mem.out


.PHONY: all build run clean info

