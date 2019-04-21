// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import (
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func Test_jwtFromHeader(t *testing.T) {
	type args struct {
		c echo.ContextInterface
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jwtFromHeader(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtFromHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jwtFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtFromQuery(t *testing.T) {
	type args struct {
		c echo.ContextInterface
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jwtFromQuery(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtFromQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jwtFromQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtFromCookie(t *testing.T) {
	type args struct {
		c echo.ContextInterface
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jwtFromCookie(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtFromCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jwtFromCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
