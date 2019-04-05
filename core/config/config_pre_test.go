// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !dev
// +build pre
// +build !prod

package config

import (
	"reflect"
	"testing"
)

func TestSetDefaults(t *testing.T) {
	type args struct {
		env map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDefaults(tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Setup()
		})
	}
}
