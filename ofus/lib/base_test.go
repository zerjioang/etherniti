// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package lib

import (
	"bytes"
	"testing"
)

func Test_encode(t *testing.T) {
	type args struct {
		nb   uint64
		buf  *bytes.Buffer
		base string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encode(tt.args.nb, tt.args.buf, tt.args.base)
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		enc  string
		base string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.enc, tt.args.base); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
