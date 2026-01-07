# 🦉 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](./_icons/coverage.svg)
![Test](https://github.com/FishGoddess/goes/actions/workflows/test.yml/badge.svg)

**Goes** is a easy-to-use and lightweight lib for executing async tasks.

[阅读中文版的文档](./README.md)

### 🥇 Features

* Limits the number of simultaneous goroutines and not reuses them by Limiter.
* Limits the number of simultaneous goroutines and reuses them by Executor.
* Supports multiple scheduling strategies, including round robin, random, etc.
* Supports spin lock with backoff strategy.
* Supports getting the number of workers available in the executor.
* Supports dynamic scaling of workers in the executor.

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 How to use

```bash
$ go get -u github.com/FishGoddess/goes
```

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FishGoddess/goes"
)

func main() {
	ctx := context.Background()

	// Limits the number of simultaneous goroutines and not reuses them.
	limiter := goes.NewLimiter(4)

	for i := 0; i < 20; i++ {
		limiter.Go(func() {
			fmt.Printf("limiter --> %s\n", time.Now())
			time.Sleep(time.Second)
		})
	}

	limiter.Wait()

	// Limits the number of simultaneous goroutines and reuses them.
	executor := goes.NewExecutor(4)
	defer executor.Close()

	for i := 0; i < 20; i++ {
		executor.Submit(ctx, func() {
			fmt.Printf("executor --> %s\n", time.Now())
			time.Sleep(time.Second)
		})
	}
}
```

_Check more examples in [_examples](./_examples)._

### 🔨 Benchmarks

```bash
$ make bench
```

```bash
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) CPU E5-26xx v4

BenchmarkLimiter-2               1256862               870.5 ns/op            24 B/op          1 allocs/op
BenchmarkExecutor-2              3916312               286.8 ns/op             0 B/op          0 allocs/op
BenchmarkAntsPool-2              1396972               846.6 ns/op             0 B/op          0 allocs/op
BenchmarkConcPool-2              1473289               843.4 ns/op             0 B/op          0 allocs/op

BenchmarkLimiterTime-2:  num is 500000, cost is 391.462505ms
BenchmarkExecutorTime-2: num is 500000, cost is 180.279155ms
BenchmarkAntsPoolTime-2: num is 500000, cost is 547.328528ms
BenchmarkConcPoolTime-2: num is 500000, cost is 390.354196ms
```

> Obviously, goes.Executor is faster than other concurrent libraries, so try goes.Executor if you prefer a light-weight and faster executor.

> Benchmarks: [_examples/basic_test.go](./_examples/basic_test.go).

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.
