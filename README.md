# 🦉 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](./_icons/coverage.svg)
![Test](https://github.com/FishGoddess/goes/actions/workflows/test.yml/badge.svg)

**Goes** 是一个简单易用的并发执行库。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 支持限制同时执行的协程数，使用 Limiter。
* 支持复用同时执行的协程数，使用 Executor。

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 使用方式

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

	// Limits goroutines.
	limiter := goes.NewLimiter(4)

	for i := 0; i < 20; i++ {
		limiter.Go(func() {
			fmt.Printf("limiter --> %s\n", time.Now())
			time.Sleep(time.Second)
		})
	}

	limiter.Wait()

	// Reuses goroutines.
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

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔨 性能测试

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

> 很明显，goes.Executor 的性能比其他的并发执行库高很多，所以当你需要一个轻量且高性能的并发执行器时，可以尝试下 goes.Executor。

> 测试文件：[_examples/basic_test.go](./_examples/basic_test.go)。

### 👥 贡献者

如果您觉得 goes 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。
