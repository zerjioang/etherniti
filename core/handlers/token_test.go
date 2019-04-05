// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/eth"
)

func TestNewTokenController(t *testing.T) {
	type args struct {
		manager eth.WalletManager
	}
	tests := []struct {
		name string
		args args
		want TokenController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenController(tt.args.manager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenController_instantiate(t *testing.T) {
	type fields struct {
		walletManager eth.WalletManager
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := TokenController{
				walletManager: tt.fields.walletManager,
			}
			if err := ctl.instantiate(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("TokenController.instantiate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTokenController_summary(t *testing.T) {
	type fields struct {
		walletManager eth.WalletManager
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := TokenController{
				walletManager: tt.fields.walletManager,
			}
			if err := ctl.summary(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("TokenController.summary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTokenController_RegisterRouters(t *testing.T) {
	type fields struct {
		walletManager eth.WalletManager
	}
	type args struct {
		router *echo.Group
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := TokenController{
				walletManager: tt.fields.walletManager,
			}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
