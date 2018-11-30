.PHONY: hello integers all iteration benchmark

default: all

all:
	go test -v ./...

hello:
	go test -v ./hello

integers:
	go test -v ./integers

iteration:
	go test -v ./iteration

benchmark:
	cd iteration && go test -v -bench=.
