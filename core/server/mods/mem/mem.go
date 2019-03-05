// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mem

import (
	"runtime"
	"sync"
	"time"

	"github.com/zerjioang/etherniti/shared/protocol"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var (
	//shared coordinator for all mem reader refresh
	ticker = time.NewTicker(5 * time.Second)
)

type MemStatus struct {
	//mem stats data holder
	m runtime.MemStats
	// locker for concurrent access
	lock *sync.Mutex
}

// constructor like function
func MemStatusMonitor() MemStatus {
	mem := MemStatus{}
	mem.lock = new(sync.Mutex)
	return mem
}

// disk usage of path/disk
func (mem *MemStatus) Start() {
	go mem.monitor()
}

func (mem *MemStatus) Read(wrapper protocol.ServerStatusResponse) protocol.ServerStatusResponse {
	mem.lock.Lock()
	wrapper.Memory.Alloc = mem.m.Alloc
	wrapper.Memory.Total = mem.m.TotalAlloc
	wrapper.Memory.Sys = mem.m.Sys
	wrapper.Memory.Mallocs = mem.m.Mallocs
	wrapper.Memory.Frees = mem.m.Frees
	wrapper.Memory.Heapalloc = mem.m.HeapAlloc

	wrapper.Gc.Numgc = mem.m.NumGC
	wrapper.Gc.NumForcedGC = mem.m.NumForcedGC
	mem.lock.Unlock()
	return wrapper
}

func (mem *MemStatus) ReadPtr(wrapper *protocol.ServerStatusResponse) {
	mem.lock.Lock()
	wrapper.Memory.Alloc = mem.m.Alloc
	wrapper.Memory.Total = mem.m.TotalAlloc
	wrapper.Memory.Sys = mem.m.Sys
	wrapper.Memory.Mallocs = mem.m.Mallocs
	wrapper.Memory.Frees = mem.m.Frees
	wrapper.Memory.Heapalloc = mem.m.HeapAlloc

	wrapper.Gc.Numgc = mem.m.NumGC
	wrapper.Gc.NumForcedGC = mem.m.NumForcedGC
	mem.lock.Unlock()
}

// internal ticker based monitor
func (mem *MemStatus) monitor() {
	for range ticker.C {
		//update latest reading information
		mem.lock.Lock()
		runtime.ReadMemStats(&mem.m)
		mem.lock.Unlock()
	}
}
