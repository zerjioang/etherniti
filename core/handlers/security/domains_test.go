// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"reflect"
	"testing"
)

func TestDomainBlacklist(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainBlacklist(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainBlacklist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainBlacklistBytesData(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainBlacklistBytesData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainBlacklistBytesData() = %v, want %v", got, tt.want)
			}
		})
	}
}
