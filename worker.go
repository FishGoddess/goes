// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type worker struct {
	funcs    chan func()
	recovery func(r any)
}

func newWorker(queueSize uint, recovery func(r any)) *worker {
	worker := &worker{
		funcs:    make(chan func(), queueSize),
		recovery: recovery,
	}

	return worker
}

func (w *worker) do(f func()) {
	if w.recovery != nil {
		defer func() {
			if r := recover(); r != nil {
				w.recovery(r)
			}
		}()
	}

	f()
}

func (w *worker) start() {
	for f := range w.funcs {
		if f == nil {
			return
		}

		w.do(f)
	}
}

func (w *worker) stop() {
	w.funcs <- nil
}

func (w *worker) submit(f func()) {
	w.funcs <- f
}
