// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mem

import (
	"github.com/zerjioang/etherniti/core/logger"
	"runtime"
	"sync"
	"sync/atomic"
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
	//shared coordinator for all mem reader instances refresh
	ticker = time.NewTicker(5 * time.Second)
	//backend default memory monitor
	mon MemStatus
)

func init(){
	mon = memStatusMonitor()
}

type MemStatus struct {
	//mem stats data holder
	m *runtime.MemStats
	// locker for concurrent access
	lock *sync.Mutex
	//flag indicating whether is running or not
	monitoring atomic.Value
}

// constructor like function
func MemStatusMonitorPtr() *MemStatus {
	return &mon
}

func MemStatusMonitor() MemStatus {
	return mon
}
// constructor like function
func memStatusMonitor() MemStatus {
	mem := MemStatus{}
	mem.lock = new(sync.Mutex)
	mem.m = new(runtime.MemStats)
	mem.monitoring.Store(false)
	return mem
}

// disk usage of path/disk
func (mem *MemStatus) Start() {
	go mem.monitor()
}

func (mem *MemStatus) Read(wrapper protocol.ServerStatusResponse) protocol.ServerStatusResponse {
	mem.lock.Lock()
	m := *mem.m
	wrapper.Memory.Alloc = m.Alloc
	wrapper.Memory.Total = m.TotalAlloc
	wrapper.Memory.Sys = m.Sys
	wrapper.Memory.Mallocs = m.Mallocs
	wrapper.Memory.Frees = m.Frees
	wrapper.Memory.Heapalloc = m.HeapAlloc

	wrapper.Gc.Numgc = m.NumGC
	wrapper.Gc.NumForcedGC = m.NumForcedGC
	mem.lock.Unlock()
	return wrapper
}

func (mem *MemStatus) ReadPtr(wrapper *protocol.ServerStatusResponse) {
	mem.lock.Lock()
	m := *mem.m
	wrapper.Memory.Alloc = m.Alloc
	wrapper.Memory.Total = m.TotalAlloc
	wrapper.Memory.Sys = m.Sys
	wrapper.Memory.Mallocs = m.Mallocs
	wrapper.Memory.Frees = m.Frees
	wrapper.Memory.Heapalloc = m.HeapAlloc

	wrapper.Gc.Numgc = m.NumGC
	wrapper.Gc.NumForcedGC = m.NumForcedGC
	mem.lock.Unlock()
}

// internal ticker based monitor
func (mem *MemStatus) monitor() {
	for range ticker.C {
		//update latest reading information
		//mem.lock.Lock()
		logger.Debug("reading node memory statistics")
		runtime.ReadMemStats(mem.m)
		//mem.lock.Unlock()
	}
}
