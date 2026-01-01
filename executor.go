// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"context"
	"errors"

	"github.com/FishGoddess/rego"
)

var ErrExecutorClosed = errors.New("goes: executor is closed")

func newExecutorClosedErr(ctx context.Context) error {
	return ErrExecutorClosed
}

type Executor struct {
	conf *config
	pool *rego.Pool[*worker]
}

func NewExecutor(workerNum uint, opts ...Option) *Executor {
	conf := newConfig().apply(opts...)

	executor := new(Executor)
	executor.conf = conf
	executor.pool = rego.New(uint64(workerNum), executor.acquire, executor.release)
	executor.pool.WithPoolClosedErrFunc(newExecutorClosedErr)
	return executor
}

func (e *Executor) acquire(ctx context.Context) (*worker, error) {
	worker := newWorker(e.conf.queueSize, e.conf.recovery)
	worker.start()
	return worker, nil
}

func (e *Executor) release(ctx context.Context, worker *worker) error {
	worker.stop()
	return nil
}

func (e *Executor) Submit(ctx context.Context, f func()) error {
	worker, err := e.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	worker.submit(f)
	return nil
}

func (e *Executor) Close(ctx context.Context) error {
	return e.pool.Close(ctx)
}
