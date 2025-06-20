# 🦉 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](./_icons/coverage.svg)
![Test](https://github.com/FishGoddess/goes/actions/workflows/test.yml/badge.svg)

**Goes** 是一个简单易用且轻量的协程数限制库。

[Read me in English](./README.en.md)

### 🥇 功能特性

* Limiter 只限制同时执行的协程数，不复用协程。

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 使用方式

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

### 👥 贡献者

如果您觉得 goes 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。
