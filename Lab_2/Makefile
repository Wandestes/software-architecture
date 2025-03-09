.PHONY: test build deps

deps:
	go mod tidy

test:
	go test ./... -v

build:
	go build -o bin/myapp Lab_2/cmd/example
