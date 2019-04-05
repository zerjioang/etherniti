// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package encoding

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewEncoding(t *testing.T) {
	type args struct {
		alphabet string
	}
	tests := []struct {
		name string
		args args
		want *Encoding
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEncoding(tt.args.alphabet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAlphabetMap(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[byte]*big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAlphabetMap(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAlphabetMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoding_Random(t *testing.T) {
	type fields struct {
		alphabet string
		index    map[byte]*big.Int
		base     *big.Int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &Encoding{
				alphabet: tt.fields.alphabet,
				index:    tt.fields.index,
				base:     tt.fields.base,
			}
			got, err := enc.Random(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoding.Random() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encoding.Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoding_MustRandom(t *testing.T) {
	type fields struct {
		alphabet string
		index    map[byte]*big.Int
		base     *big.Int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &Encoding{
				alphabet: tt.fields.alphabet,
				index:    tt.fields.index,
				base:     tt.fields.base,
			}
			if got := enc.MustRandom(tt.args.n); got != tt.want {
				t.Errorf("Encoding.MustRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoding_Base(t *testing.T) {
	type fields struct {
		alphabet string
		index    map[byte]*big.Int
		base     *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &Encoding{
				alphabet: tt.fields.alphabet,
				index:    tt.fields.index,
				base:     tt.fields.base,
			}
			if got := enc.Base(); got != tt.want {
				t.Errorf("Encoding.Base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoding_EncodeToString(t *testing.T) {
	type fields struct {
		alphabet string
		index    map[byte]*big.Int
		base     *big.Int
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &Encoding{
				alphabet: tt.fields.alphabet,
				index:    tt.fields.index,
				base:     tt.fields.base,
			}
			if got := enc.EncodeToString(tt.args.b); got != tt.want {
				t.Errorf("Encoding.EncodeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoding_DecodeString(t *testing.T) {
	type fields struct {
		alphabet string
		index    map[byte]*big.Int
		base     *big.Int
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &Encoding{
				alphabet: tt.fields.alphabet,
				index:    tt.fields.index,
				base:     tt.fields.base,
			}
			got, err := enc.DecodeString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoding.DecodeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoding.DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoding_DecodeStringN(t *testing.T) {
	type fields struct {
		alphabet string
		index    map[byte]*big.Int
		base     *big.Int
	}
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &Encoding{
				alphabet: tt.fields.alphabet,
				index:    tt.fields.index,
				base:     tt.fields.base,
			}
			got, err := enc.DecodeStringN(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoding.DecodeStringN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoding.DecodeStringN() = %v, want %v", got, tt.want)
			}
		})
	}
}
