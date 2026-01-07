// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type Task func()

func (t Task) Do(recovery func(r any)) {
	if t == nil {
		return
	}

	if recovery != nil {
		defer func() {
			if r := recover(); r != nil {
				recovery(r)
			}
		}()
	}

	t()
}
