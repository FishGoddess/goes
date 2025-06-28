// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type scheduler interface {
	// Set sets the workers to scheduler.
	Set(workers []*worker)

	// Get gets a worker from scheduler.
	Get() *worker
}

type worker struct {
	executor  *Executor
	taskQueue chan Task
}

func newWorker(executor *Executor) *worker {
	w := &worker{
		executor:  executor,
		taskQueue: make(chan Task, executor.conf.workerQueueSize),
	}

	w.work()
	return w
}

func (w *worker) handle(task Task) {
	defer func() {
		if r := recover(); r != nil {
			w.executor.conf.recover(r)
		}
	}()

	task()
}

func (w *worker) work() {
	w.executor.wg.Add(1)
	go func() {
		defer w.executor.wg.Done()
		defer close(w.taskQueue)

		for task := range w.taskQueue {
			if task == nil {
				break
			}

			w.handle(task)
		}
	}()
}

// Accept accepts a task to be handled.
func (w *worker) Accept(task Task) {
	if task == nil {
		return
	}

	w.taskQueue <- task
}

// Done signals the worker to stop working.
func (w *worker) Done() {
	w.taskQueue <- nil
}
