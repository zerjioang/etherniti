// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package disk

import (
	"sync"
	"syscall"
	"time"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var (
	//shared coordinator for all disk status reader
	ticker = time.NewTicker(5 * time.Second)
)

type DiskStatus struct {
	all        uint64
	used       uint64
	free       uint64
	monitoring bool
	lock       sync.Mutex
}

// constructor like function
func DiskUsage() DiskStatus {
	disk := DiskStatus{}
	disk.lock = sync.Mutex{}
	return disk
}

func DiskUsagePtr() *DiskStatus {
	d := DiskUsage()
	return &d
}

// disk usage of path/disk
func (disk *DiskStatus) Eval(path string) error {
	if !disk.monitoring {
		// initialize read values
		disk.read(path)
		// start monitor
		go disk.monitor(path)
	}
	return nil
}

// internal ticker based monitor
func (disk *DiskStatus) monitor(path string) error {
	disk.monitoring = true
	for !disk.monitoring {
		for range ticker.C {
			disk.read(path)
		}
	}
	return nil
}

func (disk *DiskStatus) read(path string) (*DiskStatus, error) {
	disk.lock.Lock()
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return disk, err
	}
	disk.all = fs.Blocks * uint64(fs.Bsize)
	disk.free = fs.Bfree * uint64(fs.Bsize)
	disk.used = disk.all - disk.free
	disk.lock.Unlock()
	return disk, nil
}

// get all value
func (disk DiskStatus) All() uint64 {
	disk.lock.Lock()
	defer disk.lock.Unlock()
	return disk.all / GB
}

// get used value
func (disk DiskStatus) Used() uint64 {
	disk.lock.Lock()
	defer disk.lock.Unlock()
	return disk.used / GB
}

// get free value
func (disk DiskStatus) Free() uint64 {
	disk.lock.Lock()
	defer disk.lock.Unlock()
	return disk.free / GB
}

func (disk DiskStatus) IsMonitoring() bool {
	return disk.monitoring
}
