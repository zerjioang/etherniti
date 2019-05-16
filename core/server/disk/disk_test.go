// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package disk_test

import (
	"fmt"
	"testing"

	d "github.com/zerjioang/etherniti/core/server/disk"
)

func TestDiskUsage(t *testing.T) {
	t.Run("is-monitoring-once", func(t *testing.T) {
		disk := d.DiskUsage()
		t.Log(disk.IsMonitoring())
	})
	t.Run("is-monitoring-twice", func(t *testing.T) {
		disk := d.DiskUsage()
		t.Log(disk.IsMonitoring())
		t.Log(disk.IsMonitoring())
	})
	t.Run("read-once", func(t *testing.T) {
		disk := d.DiskUsage()
		disk.Start("/")
		fmt.Printf("all: %.2f GB\n", float64(disk.All())/float64(d.GB))
		fmt.Printf("used: %.2f GB\n", float64(disk.Used())/float64(d.GB))
		fmt.Printf("free: %.2f GB\n", float64(disk.Free())/float64(d.GB))
	})
}
