// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
