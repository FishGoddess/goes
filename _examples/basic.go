// Copyright 2025s FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FishGoddess/goes"
)

func main() {
	ctx := context.Background()

	// 	// Limits the number of simultaneous goroutines and not reuses them.
	// 	limiter := goes.NewLimiter(4)
	//
	// 	for i := 0; i < 20; i++ {
	// 		limiter.Go(func() {
	// 			fmt.Printf("limiter --> %s\n", time.Now())
	// 			time.Sleep(time.Second)
	// 		})
	// 	}
	//
	// 	limiter.Wait()

	// Limits the number of simultaneous goroutines and reuses them.
	executor := goes.NewExecutor(4)
	defer executor.Close(ctx)

	for i := 0; i < 20; i++ {
		executor.Submit(ctx, func() {
			fmt.Printf("executor --> %s\n", time.Now())
			time.Sleep(time.Second)
		})
	}
}
