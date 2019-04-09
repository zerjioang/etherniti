// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package clientcache

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestCached(t *testing.T) {
	type args struct {
		c          echo.Context
		cacheValid bool
		seconds    uint
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 echo.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Cached(tt.args.c, tt.args.cacheValid, tt.args.seconds)
			if got != tt.want {
				t.Errorf("Cached() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Cached() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCachedHtml(t *testing.T) {
	type args struct {
		c           echo.Context
		cacheValid  bool
		seconds     uint
		htmlContent []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CachedHtml(tt.args.c, tt.args.cacheValid, tt.args.seconds, tt.args.htmlContent); (err != nil) != tt.wantErr {
				t.Errorf("CachedHtml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCachedJsonBlob(t *testing.T) {
	type args struct {
		c          echo.Context
		cacheValid bool
		seconds    uint
		data       []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CachedJsonBlob(tt.args.c, tt.args.cacheValid, tt.args.seconds, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CachedJsonBlob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
