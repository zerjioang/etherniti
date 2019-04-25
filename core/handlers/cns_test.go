// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/core/modules/cns"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestNewContractNameSpaceController(t *testing.T) {
	tests := []struct {
		name string
		want ContractNameSpaceController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContractNameSpaceController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContractNameSpaceController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContractNameSpaceController_register(t *testing.T) {
	type fields struct {
		ns cns.ContractNameSystem
	}
	type args struct {
		c echo.ContextInterface
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
			ctl := &ContractNameSpaceController{
				ns: tt.fields.ns,
			}
			if err := ctl.register(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ContractNameSpaceController.register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContractNameSpaceController_unregister(t *testing.T) {
	type fields struct {
		ns cns.ContractNameSystem
	}
	type args struct {
		c echo.ContextInterface
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
			ctl := &ContractNameSpaceController{
				ns: tt.fields.ns,
			}
			if err := ctl.unregister(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ContractNameSpaceController.unregister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContractNameSpaceController_resolve(t *testing.T) {
	type fields struct {
		ns cns.ContractNameSystem
	}
	type args struct {
		c echo.ContextInterface
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
			ctl := ContractNameSpaceController{
				ns: tt.fields.ns,
			}
			if err := ctl.resolve(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ContractNameSpaceController.resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContractNameSpaceController_RegisterRouters(t *testing.T) {
	type fields struct {
		ns cns.ContractNameSystem
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
			ctl := ContractNameSpaceController{
				ns: tt.fields.ns,
			}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
