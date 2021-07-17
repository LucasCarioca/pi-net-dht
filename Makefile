.PHONY := all

start:
	go run .

install:
	go get -d ./...

build-arm6:
	go build .