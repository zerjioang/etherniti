// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !pre
// +build !prod

package config

import "testing"

func TestSetDefaults(t *testing.T) {
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
			SetDefaults(tt.args.env)
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
