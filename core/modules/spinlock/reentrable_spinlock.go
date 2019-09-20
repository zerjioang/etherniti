// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package spinlock

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

// Reentrant allowable spin locks
type spinLock struct {
	owner int
	count int
}

func (sl *spinLock) Lock() {
	me := GetGoroutineId()
	if sl.owner == me {/// If the current thread has acquired the lock, the number of threads increases by one, and then returns
		sl.count++
		return
	}
	// If the lock is not acquired, spin through CAS
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}
func (sl *spinLock) Unlock() {
	if sl.owner != GetGoroutineId() {
		panic("illegalMonitorStateError")
	}
	if sl. count > 0 {// if greater than 0, it means that the current thread has acquired the lock many times, and the release lock is simulated by subtracting count from one.
		sl.count--
	} else {
		// If count== 0, the lock can be released, which ensures that the number of acquisitions of the lock is the same as the number of releases of the lock.
		atomic.StoreUint32((*uint32)(sl), 0)
	}
}

func GetGoroutineId() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func NewSpinLock() sync.Locker {
	var lock spinLock
	return &lock
}

