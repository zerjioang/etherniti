// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"testing"
)

func TestIpToUint32(t *testing.T) {

	t.Run("convert-bytes", func(t *testing.T) {
		intVal := Ip2int("101.41.132.176")
		t.Log("uint32 ip:", intVal)
		if intVal != 1697219760 {
			t.Error("failed to convert ip to numeric")
		}
	})
	t.Run("convert-uint32", func(t *testing.T) {
		ipStr := Int2ip(1697219760)
		t.Log("str ip:", string(ipStr))
		if string(ipStr) != "101.41.132.176" {
			t.Error("failed to convert ip to numeric")
		}
	})
}

func TestIp2int(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ip2int(tt.args.ip); got != tt.want {
				t.Errorf("Ip2int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt2ip(t *testing.T) {
	type args struct {
		ipInt int64
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
			if got := Int2ip(tt.args.ipInt); got != tt.want {
				t.Errorf("Int2ip() = %v, want %v", got, tt.want)
			}
		})
	}
}
