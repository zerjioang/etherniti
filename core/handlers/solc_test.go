// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestNewSolcController(t *testing.T) {
	tests := []struct {
		name string
		want SolcController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSolcController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSolcController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolcController_version(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := &SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.version(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.version() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_compileSingle(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.compileSingle(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.compileSingle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_compileSingleFromBase64(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.compileSingleFromBase64(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.compileSingleFromBase64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_compileMultiple(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.compileMultiple(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.compileMultiple() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_compileFromGit(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.compileFromGit(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.compileFromGit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_compileFromUploadedZip(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.compileFromUploadedZip(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.compileFromUploadedZip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_compileFromUploadedTargz(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.compileFromUploadedTargz(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.compileFromUploadedTargz() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_compileModeSelector(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			if err := ctl.compileModeSelector(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SolcController.compileModeSelector() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSolcController_RegisterRouters(t *testing.T) {
	type fields struct {
		solidityVersionResponse atomic.Value
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
			ctl := SolcController{
				solidityVersionResponse: tt.fields.solidityVersionResponse,
			}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
