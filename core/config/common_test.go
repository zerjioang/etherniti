// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

func TestRead(t *testing.T) {
	type args struct {
		env map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Read(tt.args.env)
		})
	}
}

func TestGetRedirectUrl(t *testing.T) {
	type args struct {
		host string
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"redirect",
			args{"127.0.0.1", "/test"},
			"https://127.0.0.1:8080/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedirectUrl(tt.args.host, tt.args.path); got != tt.want {
				t.Errorf("GetRedirectUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCertPem(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCertPem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCertPem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKeyPem(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKeyPem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeyPem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHttpMode(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHttpMode(); got != tt.want {
				t.Errorf("IsHttpMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSocketMode(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSocketMode(); got != tt.want {
				t.Errorf("IsSocketMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsProfilingEnabled(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProfilingEnabled(); got != tt.want {
				t.Errorf("IsProfilingEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceListeningMode(t *testing.T) {
	tests := []struct {
		name string
		want listener.ServiceType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ServiceListeningMode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceListeningMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
