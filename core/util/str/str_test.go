// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package str

import (
	"reflect"
	"testing"
)

func TestGetJsonBytes(t *testing.T) {
	t.Run("get-bytes-nil", func(t *testing.T) {
		GetJsonBytes(nil)
	})
}

func TestToLowerAscii(t *testing.T) {
	t.Run("ToLowerAscii", func(t *testing.T) {
		val := "Hello World, This is AWESOME"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		if converted != "hello world, this is awesome" {
			t.Error("failed to lowercase")
		}
	})
	t.Run("ToLowerAscii-ua", func(t *testing.T) {
		val := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:61.0) Gecko/20100101 Firefox/61.0"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		if converted != "mozilla/5.0 (x11; ubuntu; linux x86_64; rv:61.0) gecko/20100101 firefox/61.0" {
			t.Error("failed to lowercase")
		}
	})
}

func TestBytes(t *testing.T) {
	type args struct {
		data string
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
			if got := UnsafeBytes(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsafeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		data []byte
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
			if got := UnsafeString(tt.args.data); got != tt.want {
				t.Errorf("UnsafeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStdMarshal(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdMarshal(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("StdMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StdMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
