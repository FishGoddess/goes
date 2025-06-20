# 🦉 Goes

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/goes)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](./_icons/coverage.svg)
![Test](https://github.com/FishGoddess/goes/actions/workflows/test.yml/badge.svg)

**Goes** is a easy-to-use and lightweight lib for limiting goroutines.

[阅读中文版的文档](./README.md)

### 🥇 Features

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

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.
