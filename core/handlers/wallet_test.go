// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/shared/protocol"
)

func TestNewWalletController(t *testing.T) {
	tests := []struct {
		name string
		want WalletController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWalletController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWalletController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletController_Mnemonic(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     WalletController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			if err := ctl.Mnemonic(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("WalletController.Mnemonic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWalletController_HdWallet(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     WalletController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			if err := ctl.HdWallet(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("WalletController.HdWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWalletController_Entropy(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     WalletController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			if err := ctl.Entropy(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("WalletController.Entropy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWalletController_generateSecureEntropy(t *testing.T) {
	type args struct {
		request protocol.EntropyRequest
	}
	tests := []struct {
		name    string
		ctl     WalletController
		args    args
		want    protocol.EntropyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			got, err := ctl.generateSecureEntropy(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("WalletController.generateSecureEntropy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalletController.generateSecureEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletController_createHdWallet(t *testing.T) {
	type args struct {
		request protocol.NewHdWalletRequest
	}
	tests := []struct {
		name    string
		ctl     WalletController
		args    args
		want    protocol.HdWalletResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			got, err := ctl.createHdWallet(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("WalletController.createHdWallet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalletController.createHdWallet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletController_generateAddress(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     WalletController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			if err := ctl.generateAddress(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("WalletController.generateAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWalletController_isValidAddress(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     WalletController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			if err := ctl.isValidAddress(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("WalletController.isValidAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWalletController_RegisterRouters(t *testing.T) {
	type args struct {
		router *echo.Group
	}
	tests := []struct {
		name string
		ctl  WalletController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}

func TestWalletController_getIntParam(t *testing.T) {
	type args struct {
		c   echo.Context
		key string
	}
	tests := []struct {
		name string
		ctl  WalletController
		args args
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := WalletController{}
			if got := ctl.getIntParam(tt.args.c, tt.args.key); got != tt.want {
				t.Errorf("WalletController.getIntParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
