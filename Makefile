.PHONY := all

start:
	go run .

install:
	go get -d ./...

build:
	go build .
