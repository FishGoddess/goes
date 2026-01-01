.PHONY: test fmt

all: test

test:
	go test -v -cover ./...

bench:
	go test -v ./_examples/basic_test.go -run=none -bench=. -benchmem -benchtime=1s

fmt:
	go fmt ./...
