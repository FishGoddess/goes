.PHONY: test fmt

all: test

test:
	go test -v -cover ./...

bench:
	go test -v ./_examples/performance_test.go -run=none -bench=. -benchmem -benchtime=1s

fmt:
	go fmt ./...
