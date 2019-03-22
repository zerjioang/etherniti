// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package spinlock

import (
	"sync/atomic"
)

type SpinLock struct {
	state int32
}

func (s *SpinLock) TyLock() bool {
	return atomic.CompareAndSwapInt32(&s.state, 0, 1)
}

func (s *SpinLock) lock() bool {
	return atomic.CompareAndSwapInt32(&s.state, 0, 1)
}

func (s SpinLock) IsLocked() bool {
	return s.state == 1
}

func (s *SpinLock) Unlock() {
	s.state = 0
}
