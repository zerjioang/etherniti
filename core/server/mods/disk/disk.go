package disk

import (
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
	// ticker = time.NewTicker(5 * time.Second)
	ticker = time.NewTicker(100 * time.Millisecond)
)

type DiskStatus struct {
	all  uint64
	used uint64
	free uint64
	//monitoring bool
	monitoring bool
}

// constructor like function
func DiskUsage() DiskStatus {
	disk := DiskStatus{}
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
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return disk, err
	}
	disk.all = fs.Blocks * uint64(fs.Bsize)
	disk.free = fs.Bfree * uint64(fs.Bsize)
	disk.used = disk.all - disk.free
	return disk, nil
}

// get all value
func (disk DiskStatus) All() uint64 {
	return disk.all
}

// get used value
func (disk DiskStatus) Used() uint64 {
	return disk.used
}

// get free value
func (disk DiskStatus) Free() uint64 {
	return disk.free
}

func (disk DiskStatus) IsMonitoring() bool {
	return disk.monitoring
}
