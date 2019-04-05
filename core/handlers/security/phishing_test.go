// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"reflect"
	"testing"
)

func TestPhishingBlacklistRawBytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhishingBlacklistRawBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PhishingBlacklistRawBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhishingWhitelistRawBytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhishingWhitelistRawBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PhishingWhitelistRawBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuzzyDataRawBytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuzzyDataRawBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FuzzyDataRawBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		arr []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.arr, tt.args.str); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDangerousDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDangerousDomain(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsDangerousDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
