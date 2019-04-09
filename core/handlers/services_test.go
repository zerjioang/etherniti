// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func Test_infuraJwt(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := infuraJwt(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("infuraJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quiknodeJwt(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quiknodeJwt(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quiknodeJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_privateJwt(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := privateJwt(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("privateJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwt(t *testing.T) {
	type args struct {
		next     echo.HandlerFunc
		errorMsg string
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jwt(tt.args.next, tt.args.errorMsg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_next(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := next(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterServices(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
		want *echo.Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterServices(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterServices() = %v, want %v", got, tt.want)
			}
		})
	}
}
