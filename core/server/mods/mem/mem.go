// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mem

import (
	"runtime"
	"sync/atomic"
	"time"

	"github.com/zerjioang/etherniti/core/logger"

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

func init() {
	mon = memStatusMonitor()
}

type MemStatus struct {
	//mem stats data holder
	m runtime.MemStats
	//flag indicating whether is running or not
	monitoring atomic.Value
}

// constructor like function as ptr
func MemStatusMonitorPtr() *MemStatus {
	return &mon
}

// constructor like function as struct
func MemStatusMonitor() MemStatus {
	return mon
}

// internal used constructor like function
func memStatusMonitor() MemStatus {
	mem := MemStatus{}
	mem.monitoring.Store(false)
	return mem
}

// starts a background routine checking for memory status
func (mem *MemStatus) Start() {
	isRunning := mem.monitoring.Load().(bool)
	if !isRunning {
		mem.monitoring.Store(true)
		go mem.monitor()
	}
}

func (mem MemStatus) Read(wrapper protocol.ServerStatusResponse) protocol.ServerStatusResponse {
	wrapper.Memory.Alloc = mem.m.Alloc
	wrapper.Memory.Total = mem.m.TotalAlloc
	wrapper.Memory.Sys = mem.m.Sys
	wrapper.Memory.Mallocs = mem.m.Mallocs
	wrapper.Memory.Frees = mem.m.Frees
	wrapper.Memory.Heapalloc = mem.m.HeapAlloc

	wrapper.Gc.Numgc = mem.m.NumGC
	wrapper.Gc.NumForcedGC = mem.m.NumForcedGC
	return wrapper
}

func (mem MemStatus) ReadPtr(wrapper *protocol.ServerStatusResponse) {
	wrapper.Memory.Alloc = mem.m.Alloc
	wrapper.Memory.Total = mem.m.TotalAlloc
	wrapper.Memory.Sys = mem.m.Sys
	wrapper.Memory.Mallocs = mem.m.Mallocs
	wrapper.Memory.Frees = mem.m.Frees
	wrapper.Memory.Heapalloc = mem.m.HeapAlloc

	wrapper.Gc.Numgc = mem.m.NumGC
	wrapper.Gc.NumForcedGC = mem.m.NumForcedGC
}

// reads backend memory statistics
func (mem *MemStatus) ReadMemory() {
	logger.Debug("reading node memory statistics")
	runtime.ReadMemStats(&mem.m)
}

// internal ticker based monitor
func (mem MemStatus) monitor() {
	for range ticker.C {
		//update latest reading information
		mem.ReadMemory()
	}
}
