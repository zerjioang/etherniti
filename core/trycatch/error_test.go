// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package trycatch

import "testing"

func TestNil(t *testing.T) {
	tests := []struct {
		name string
		want Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nil(); got != tt.want {
				t.Errorf("Nil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRet(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
		want Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ret(tt.args.e); got != tt.want {
				t.Errorf("Ret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.msg); got != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	tests := []struct {
		name  string
		stack Error
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Occur(t *testing.T) {
	tests := []struct {
		name  string
		stack Error
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.Occur(); got != tt.want {
				t.Errorf("Error.Occur() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_None(t *testing.T) {
	tests := []struct {
		name  string
		stack Error
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stack.None(); got != tt.want {
				t.Errorf("Error.None() = %v, want %v", got, tt.want)
			}
		})
	}
}
