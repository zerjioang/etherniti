// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"reflect"
	"testing"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/trycatch"
)

func TestWrapper(t *testing.T) {

}

func TestSendSuccess(t *testing.T) {
	type args struct {
		c        echo.Context
		logMsg   string
		response interface{}
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
			if err := SendSuccess(tt.args.c, tt.args.logMsg, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("SendSuccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendSuccessBlob(t *testing.T) {
	type args struct {
		c   echo.Context
		raw []byte
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
			if err := SendSuccessBlob(tt.args.c, tt.args.raw); (err != nil) != tt.wantErr {
				t.Errorf("SendSuccessBlob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSuccess(t *testing.T) {
	type args struct {
		c      echo.Context
		msg    string
		result string
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
			if err := Success(tt.args.c, tt.args.msg, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Success() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToSuccess(t *testing.T) {
	type args struct {
		msg    string
		result interface{}
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
			if got := ToSuccess(tt.args.msg, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorStr(t *testing.T) {
	type args struct {
		c   echo.Context
		str string
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
			if err := ErrorStr(tt.args.c, tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("ErrorStr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		c   echo.Context
		err error
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
			if err := Error(tt.args.c, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStackError(t *testing.T) {
	type args struct {
		c        echo.Context
		stackErr trycatch.Error
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
			if err := StackError(tt.args.c, tt.args.stackErr); (err != nil) != tt.wantErr {
				t.Errorf("StackError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
