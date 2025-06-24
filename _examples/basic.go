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
