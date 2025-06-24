// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type Task = func()

type worker struct {
	executor  *Executor
	taskQueue chan Task
}

func newWorker(executor *Executor) *worker {
	taskQueue := make(chan Task, executor.conf.workerQueueSize)

	w := &worker{executor: executor, taskQueue: taskQueue}
	w.work()
	return w
}

func (w *worker) handle(task func()) {
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

		for task := range w.taskQueue {
			if task == nil {
				break
			}

			w.handle(task)
		}

		close(w.taskQueue)
	}()
}

// Accept accepts a task to be handled.
func (w *worker) Accept(task func()) {
	if task == nil {
		return
	}

	w.taskQueue <- task
}

// Done signals the worker to stop working.
func (w *worker) Done() {
	w.taskQueue <- nil
}
