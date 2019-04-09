// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestNewSecurityController(t *testing.T) {
	tests := []struct {
		name string
		want SecurityController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSecurityController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSecurityController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurityController_domainBlacklist(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     SecurityController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := SecurityController{}
			if err := ctl.domainBlacklist(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SecurityController.domainBlacklist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSecurityController_phisingWhitelist(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     SecurityController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := SecurityController{}
			if err := ctl.phisingWhitelist(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SecurityController.phisingWhitelist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSecurityController_phisingBlacklist(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     SecurityController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := SecurityController{}
			if err := ctl.phisingBlacklist(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SecurityController.phisingBlacklist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSecurityController_fuzzylist(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     SecurityController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := SecurityController{}
			if err := ctl.fuzzylist(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SecurityController.fuzzylist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSecurityController_isDangerousDomain(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     SecurityController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := SecurityController{}
			if err := ctl.isDangerousDomain(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SecurityController.isDangerousDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSecurityController_RegisterRouters(t *testing.T) {
	type args struct {
		router *echo.Group
	}
	tests := []struct {
		name string
		ctl  SecurityController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := SecurityController{}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
