.PHONY: test fmt

all: test

test:
	go test -cover ./...

fmt:
	go fmt ./...
