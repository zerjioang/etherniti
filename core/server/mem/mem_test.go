// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mem

import (
	"reflect"
	"runtime"
	"sync/atomic"
	"testing"

	"github.com/zerjioang/etherniti/shared/protocol"
)

func TestMemStatus(t *testing.T) {

}

func TestMemStatusMonitorPtr(t *testing.T) {
	tests := []struct {
		name string
		want *MemStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MemStatusMonitorPtr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemStatusMonitorPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStatusMonitor(t *testing.T) {
	tests := []struct {
		name string
		want MemStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MemStatusMonitor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemStatusMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memStatusMonitor(t *testing.T) {
	tests := []struct {
		name string
		want MemStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := memStatusMonitor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memStatusMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStatus_Start(t *testing.T) {
	type fields struct {
		m          runtime.MemStats
		monitoring atomic.Value
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := &MemStatus{
				m:          tt.fields.m,
				monitoring: tt.fields.monitoring,
			}
			mem.Start()
		})
	}
}

func TestMemStatus_Read(t *testing.T) {
	type fields struct {
		m          runtime.MemStats
		monitoring atomic.Value
	}
	type args struct {
		wrapper protocol.ServerStatusResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   protocol.ServerStatusResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := MemStatus{
				m:          tt.fields.m,
				monitoring: tt.fields.monitoring,
			}
			if got := mem.Read(tt.args.wrapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemStatus.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStatus_ReadPtr(t *testing.T) {
	type fields struct {
		m          runtime.MemStats
		monitoring atomic.Value
	}
	type args struct {
		wrapper *protocol.ServerStatusResponse
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
			mem := MemStatus{
				m:          tt.fields.m,
				monitoring: tt.fields.monitoring,
			}
			mem.ReadPtr(tt.args.wrapper)
		})
	}
}

func TestMemStatus_ReadMemory(t *testing.T) {
	type fields struct {
		m          runtime.MemStats
		monitoring atomic.Value
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := &MemStatus{
				m:          tt.fields.m,
				monitoring: tt.fields.monitoring,
			}
			mem.ReadMemory()
		})
	}
}

func TestMemStatus_monitor(t *testing.T) {
	type fields struct {
		m          runtime.MemStats
		monitoring atomic.Value
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := MemStatus{
				m:          tt.fields.m,
				monitoring: tt.fields.monitoring,
			}
			mem.monitor()
		})
	}
}
