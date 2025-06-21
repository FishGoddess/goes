// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"

	"github.com/FishGoddess/goes/pkg/spinlock"
)

func main() {
	// It's an implementation of sync.Locker, so you can use it as a mutex.
	spin := spinlock.New()

	var total int
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			spin.Lock()
			total++
			spin.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("total is %d", total)
}
