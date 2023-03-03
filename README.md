# 📝 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**Goes** 是一个比协程池轻量的协程数限制库。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 限制协程数，没了。

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

最后，我想感谢 JetBrains 公司的 **free JetBrains Open Source license(s)**，因为 goes 是用该计划下的 Idea / GoLand 完成开发的。

<a href="https://www.jetbrains.com/?from=goes" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>