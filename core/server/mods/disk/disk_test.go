// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package disk

import (
	"fmt"
	"testing"
)

func TestDiskUsage(t *testing.T) {
	t.Run("is-monitoring-once", func(t *testing.T) {
		disk := DiskUsage()
		t.Log(disk.IsMonitoring())
	})
	t.Run("is-monitoring-twice", func(t *testing.T) {
		disk := DiskUsage()
		t.Log(disk.IsMonitoring())
		t.Log(disk.IsMonitoring())
	})
	t.Run("read-once", func(t *testing.T) {
		disk := DiskUsage()
		err := disk.Start("/")
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("all: %.2f GB\n", float64(disk.All())/float64(GB))
		fmt.Printf("used: %.2f GB\n", float64(disk.Used())/float64(GB))
		fmt.Printf("free: %.2f GB\n", float64(disk.Free())/float64(GB))
	})
}
