all: build test

build:
	go generate .
	go build .
test:
	go test -json ./... | tparse -all

.PHONY: build test all
