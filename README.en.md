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

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 How to use

```bash
$ go get -u github.com/FishGoddess/goes
```

```go
package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/goes"
)

func main() {
	// Limits the number of simultaneous goroutines and not reuses them.
	limiter := goes.NewLimiter(4)

	for i := 0; i < 20; i++ {
		limiter.Go(func() {
			fmt.Println("limiter --> ", time.Now())
			time.Sleep(time.Second)
		})
	}

	limiter.Wait()

	// Limits the number of simultaneous goroutines and reuses them.
	executor := goes.NewExecutor(4)
	defer executor.Close()

	for i := 0; i < 20; i++ {
		executor.Submit(func() {
			fmt.Println("executor --> ", time.Now())
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
cpu: AMD EPYC 7K62 48-Core Processor

BenchmarkLimiter-2               2417040               498.5 ns/op            24 B/op          1 allocs/op
BenchmarkExecutor-2             20458502                58.3 ns/op             0 B/op          0 allocs/op
BenchmarkAntsPool-2              4295964               271.7 ns/op             0 B/op          0 allocs/op

BenchmarkLimiterTime-2:  num is 1000000, cost is 300.936441ms
BenchmarkExecutorTime-2: num is 1000000, cost is  63.026947ms
BenchmarkAntsPoolTime-2: num is  999744, cost is 346.972287ms
```

> Obviously, goes.Executor is 5x faster than ants.Pool which has more features, so try goes if you prefer a lightweight and faster executor.

> Benchmarks: [_examples/performance_test.go](./_examples/performance_test.go).

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.
