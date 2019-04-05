// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"

	"github.com/labstack/echo"
)

func TestNewWeb3Controller(t *testing.T) {
	tests := []struct {
		name string
		want Web3Controller
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeb3Controller(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeb3Controller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeb3Controller_getBalance(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.getBalance(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.getBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_getBalanceAtBlock(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.getBalanceAtBlock(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.getBalanceAtBlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_getNodeInfo(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.getNodeInfo(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.getNodeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_getNetworkVersion(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.getNetworkVersion(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.getNetworkVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_makeRpcCallNoParams(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.makeRpcCallNoParams(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.makeRpcCallNoParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_getAccountsWithBalance(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.getAccountsWithBalance(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.getAccountsWithBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_isContractAddress(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.isContractAddress(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.isContractAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20Name(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20Name(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20Name() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20Symbol(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20Symbol(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20Symbol() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20totalSupply(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20totalSupply(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20totalSupply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20decimals(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20decimals(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20decimals() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20Balanceof(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20Balanceof(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20Balanceof() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20Allowance(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20Allowance(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20Allowance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20Transfer(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20Transfer(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20Approve(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20Approve(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20Approve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_erc20TransferFrom(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.erc20TransferFrom(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.erc20TransferFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_deployContract(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := &Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			if err := ctl.deployContract(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Web3Controller.deployContract() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeb3Controller_RegisterRouters(t *testing.T) {
	type fields struct {
		NetworkController NetworkController
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
			ctl := Web3Controller{
				NetworkController: tt.fields.NetworkController,
			}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
