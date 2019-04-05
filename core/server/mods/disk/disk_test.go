// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package disk_test

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	d "github.com/zerjioang/etherniti/core/server/mods/disk"
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

func TestDiskUsagePtr(t *testing.T) {
	tests := []struct {
		name string
		want *DiskStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiskUsagePtr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiskUsagePtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskStatus_Start(t *testing.T) {
	type fields struct {
		all  uint64
		used uint64
		free uint64
		lock *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := &DiskStatus{
				all:  tt.fields.all,
				used: tt.fields.used,
				free: tt.fields.free,
				lock: tt.fields.lock,
			}
			disk.Start(tt.args.path)
		})
	}
}

func TestDiskStatus_monitor(t *testing.T) {
	type fields struct {
		all  uint64
		used uint64
		free uint64
		lock *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := &DiskStatus{
				all:  tt.fields.all,
				used: tt.fields.used,
				free: tt.fields.free,
				lock: tt.fields.lock,
			}
			disk.monitor(tt.args.path)
		})
	}
}

func TestDiskStatus_read(t *testing.T) {
	type fields struct {
		all  uint64
		used uint64
		free uint64
		lock *sync.Mutex
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := &DiskStatus{
				all:  tt.fields.all,
				used: tt.fields.used,
				free: tt.fields.free,
				lock: tt.fields.lock,
			}
			if err := disk.read(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("DiskStatus.read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiskStatus_All(t *testing.T) {
	type fields struct {
		all  uint64
		used uint64
		free uint64
		lock *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := &DiskStatus{
				all:  tt.fields.all,
				used: tt.fields.used,
				free: tt.fields.free,
				lock: tt.fields.lock,
			}
			if got := disk.All(); got != tt.want {
				t.Errorf("DiskStatus.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskStatus_Used(t *testing.T) {
	type fields struct {
		all  uint64
		used uint64
		free uint64
		lock *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := &DiskStatus{
				all:  tt.fields.all,
				used: tt.fields.used,
				free: tt.fields.free,
				lock: tt.fields.lock,
			}
			if got := disk.Used(); got != tt.want {
				t.Errorf("DiskStatus.Used() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskStatus_Free(t *testing.T) {
	type fields struct {
		all  uint64
		used uint64
		free uint64
		lock *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := &DiskStatus{
				all:  tt.fields.all,
				used: tt.fields.used,
				free: tt.fields.free,
				lock: tt.fields.lock,
			}
			if got := disk.Free(); got != tt.want {
				t.Errorf("DiskStatus.Free() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskStatus_IsMonitoring(t *testing.T) {
	type fields struct {
		all  uint64
		used uint64
		free uint64
		lock *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disk := DiskStatus{
				all:  tt.fields.all,
				used: tt.fields.used,
				free: tt.fields.free,
				lock: tt.fields.lock,
			}
			if got := disk.IsMonitoring(); got != tt.want {
				t.Errorf("DiskStatus.IsMonitoring() = %v, want %v", got, tt.want)
			}
		})
	}
}
