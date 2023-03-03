# 📝 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**Goes** is a lightweight lib for limiting goroutines.

[阅读中文版的文档](./README.md)

### 🥇 Features

* Limit goroutines, that's it.

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

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

At last, I want to thank JetBrains for **free JetBrains Open Source license(s)**, because goes is developed with Idea /
GoLand under it.

<a href="https://www.jetbrains.com/?from=goes" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>