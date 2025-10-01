all: build test

build:
	go build .
test:
	go test -json ./... | tparse -all

.PHONY: build test all
