# 🦉 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](./_icons/coverage.svg)
![Test](https://github.com/FishGoddess/goes/actions/workflows/test.yml/badge.svg)

**Goes** 是一个简单易用且轻量的异步任务执行库。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 支持只限制同时执行的协程数，但不复用协程，使用 Limiter。
* 支持限制同时执行的协程数，且复用协程，使用 Executor。
* 支持多种调度策略，包括轮询、随机等。
* 支持使用退避策略的自旋锁。
* 支持查询当前的 worker 数量。

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 使用方式

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

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔨 性能测试

```bash
$ make bench
```

```bash
goos: linux
goarch: amd64
cpu: AMD EPYC 7K62 48-Core Processor

BenchmarkLimiter-2               2417040               498.5 ns/op            24 B/op          1 allocs/op
BenchmarkExecutor-2             23793781                49.9 ns/op             0 B/op          0 allocs/op
BenchmarkAntsPool-2              4295964               271.7 ns/op             0 B/op          0 allocs/op

BenchmarkLimiterTime-2:  num is 1000000, cost is 300.936441ms
BenchmarkExecutorTime-2: num is 1000000, cost is  51.350509ms
BenchmarkAntsPoolTime-2: num is  999744, cost is 346.972287ms
```

> 很明显，goes.Executor 的性能比功能更丰富的 ants.Pool 要高出 5 倍左右，所以当你需要一个很轻量且高性能的异步任务执行器时，可以尝试下 goes。

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)。

### 👥 贡献者

如果您觉得 goes 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。
