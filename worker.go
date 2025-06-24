// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type worker struct {
	pool  *Pool
	tasks chan func()
}

func newWorker(pool *Pool) *worker {
	w := &worker{
		pool:  pool,
		tasks: make(chan func(), pool.conf.queueSize),
	}

	w.work()
	return w
}

func (w *worker) handle(task func()) {
	if task == nil {
		return
	}

	if w.pool.conf.recoverFunc != nil {
		defer func() {
			if r := recover(); r != nil {
				w.pool.conf.recoverFunc(r)
			}
		}()
	}

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

func (w *worker) Accept(task func()) {
	if task == nil {
		return
	}

	w.tasks <- task
}

func (w *worker) Done() {
	w.tasks <- nil
}
