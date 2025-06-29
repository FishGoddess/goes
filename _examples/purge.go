// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/FishGoddess/goes"
)

func watchExecutorWorkers(executor *goes.Executor) {
	for {
		fmt.Printf("workers --> %d\n", executor.AvailableWorkers())
		time.Sleep(time.Second)
	}
}

func main() {
	// Creates a purge-active executor with 16 workers.
	// The purge task will be executed every purge interval and purge workers not in lifetime.
	purgeInterval := time.Minute
	workerLifetime := time.Minute

	executor := goes.NewExecutor(16, goes.WithPurgeActive(purgeInterval, workerLifetime))
	defer executor.Close()

	go watchExecutorWorkers(executor)

	// Submit some tasks.
	for i := 0; i < 200; i++ {
		no := i

		executor.Submit(func() {
			r := 3000 + rand.Intn(1000)
			fmt.Printf("task %d --> %d\n", no, r)
			time.Sleep(time.Duration(r) * time.Millisecond)
		})
	}

	time.Sleep(time.Minute + 5*time.Second)
	fmt.Println("Wait! You will see the workers are decreasing and increasing...")

	for i := 0; i < 10; i++ {
		no := i

		executor.Submit(func() {
			r := 10000 + rand.Intn(1000)
			fmt.Printf("task %d --> %d\n", no, r)
			time.Sleep(time.Duration(r) * time.Millisecond)
		})
	}

	time.Sleep(2 * time.Minute)
}
