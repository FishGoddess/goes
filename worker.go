// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type worker struct {
	pool  *Pool
	tasks chan func()
}

func newWorker(pool *Pool) *worker {
	tasks := make(chan func(), pool.conf.queueSize)

	w := &worker{pool: pool, tasks: tasks}
	w.work()
	return w
}

func (w *worker) handle(task func()) {
	defer func() {
		if r := recover(); r != nil {
			w.pool.conf.recover(r)
		}
	}()

	task()
}

func (w *worker) work() {
	w.pool.wg.Add(1)
	go func() {
		defer w.pool.wg.Done()

		for task := range w.tasks {
			if task == nil {
				break
			}

			w.handle(task)
		}

		close(w.tasks)
	}()
}

// Accept accepts a task to be handled.
func (w *worker) Accept(task func()) {
	if task == nil {
		return
	}

	w.tasks <- task
}

// Done signals the worker to stop working.
func (w *worker) Done() {
	w.tasks <- nil
}
