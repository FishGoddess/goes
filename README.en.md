# 🦉 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](./_icons/coverage.svg)
![Test](https://github.com/FishGoddess/goes/actions/workflows/test.yml/badge.svg)

**Goes** is a easy-to-use and lightweight lib for limiting goroutines.

[阅读中文版的文档](./README.md)

### 🥇 Features

* Spin lock with backoff strategy.
* Limiter only limits the number of simultaneous goroutines, not reuses goroutines.

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

### 🚀 How to use

$ go get -u github.com/FishGoddess/goes

```go
package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/goes"
)

func main() {
	limiter := goes.NewLimiter(4)

	for i := 0; i < 100; i++ {
		limiter.Go(func() {
			fmt.Println(time.Now())
			time.Sleep(100 * time.Millisecond)
		})
	}

	limiter.Wait()
}
```

_Check more examples in [_examples](./_examples)._

### 🗡️ Benchmarks

```bash
$ make bench
```

```bash
goos: linux
goarch: amd64
cpu: AMD EPYC 7K62 48-Core Processor

BenchmarkLimiter-2               2417040               498.5 ns/op            24 B/op          1 allocs/op
BenchmarkPool-2                 23793781                49.9 ns/op             0 B/op          0 allocs/op
BenchmarkAntsPool-2              4295964               271.7 ns/op             0 B/op          0 allocs/op

BenchmarkLimiterTime-2:  num is 1000000, cost is 300.936441ms
BenchmarkPoolTime-2:     num is 1000000, cost is  51.350509ms
BenchmarkAntsPoolTime-2: num is  999744, cost is 346.972287ms
```

> As you can see, goes.Pool is 5x faster than ants.Pool which has more features, so try goes if you want better performance and less features.

> Benchmarks: [_examples/performance_test.go](./_examples/performance_test.go).

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.
