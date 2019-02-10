package disk

import (
	"fmt"
	"testing"
)

func TestDiskUsage(t *testing.T) {
	disk := DiskUsage("/")
	fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
}
