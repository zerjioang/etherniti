// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

import (
	"bytes"
	"reflect"
	"testing"
)

func TestServerStatusResponse_Bytes(t *testing.T) {
	type fields struct {
		Cpus    Cpus
		Runtime Runtime
		Version Version
		Disk    Disk
		Memory  Memory
		Gc      Gc
	}
	type args struct {
		buffer *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ServerStatusResponse{
				Cpus:    tt.fields.Cpus,
				Runtime: tt.fields.Runtime,
				Version: tt.fields.Version,
				Disk:    tt.fields.Disk,
				Memory:  tt.fields.Memory,
				Gc:      tt.fields.Gc,
			}
			if got := r.Bytes(tt.args.buffer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerStatusResponse.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itoa(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itoa(tt.args.v); got != tt.want {
				t.Errorf("itoa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itoau32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itoau32(tt.args.v); got != tt.want {
				t.Errorf("itoau32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itoau64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itoau64(tt.args.v); got != tt.want {
				t.Errorf("itoau64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tofloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tofloat64(tt.args.v); got != tt.want {
				t.Errorf("tofloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
