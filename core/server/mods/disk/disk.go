package disk

import (
	"fmt"
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
	//shared coorditator for all disk status reader
	ticker = time.NewTicker(5 * time.Second)
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	monitoring bool
}

// constructor like function
func DiskUsage() DiskStatus {
	disk := DiskStatus{}
	return disk
}

// disk usage of path/disk
func (disk DiskStatus) Eval(path string) (DiskStatus, error) {
	if !disk.monitoring {
		// start monitor
		go disk.monitor(path)
		disk.monitoring = true
		// initialize read values
		return disk.read(path)
	}
	return disk, nil
}

// internal ticker based monitor
func (disk DiskStatus) monitor(path string) (DiskStatus, error) {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		return disk.read(path)
	}
	return disk, nil
}

func (disk DiskStatus) read(path string) (DiskStatus, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return disk, err
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return disk, nil
}
