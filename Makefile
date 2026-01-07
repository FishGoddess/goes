.PHONY: fmt test

all: fmt test

fmt:
	go fmt ./...

test:
	go test -v -cover ./...

bench:
	go test -v ./_examples/basic_test.go -run=none -bench=. -benchmem -benchtime=1s

