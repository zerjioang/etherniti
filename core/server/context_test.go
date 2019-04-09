// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/etherniti/core/eth/profile"
	ethrpc "github.com/zerjioang/etherniti/core/eth/rpc"
)

func TestNewEthernitiContext(t *testing.T) {
	t.Run("instantiate-nil", func(t *testing.T) {
		_ = NewEthernitiContext(nil)
	})
}

func TestEthernitiContext_ConnectionProfileSetup(t *testing.T) {
	type fields struct {
		Context     echo.Context
		profileData profile.ConnectionProfile
	}
	tests := []struct {
		name    string
		fields  fields
		want    profile.ConnectionProfile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := &EthernitiContext{
				Context:     tt.fields.Context,
				profileData: tt.fields.profileData,
			}
			got, err := context.ConnectionProfileSetup()
			if (err != nil) != tt.wantErr {
				t.Errorf("EthernitiContext.ConnectionProfileSetup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EthernitiContext.ConnectionProfileSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthernitiContext_RecoverEthClientFromTokenOrPeerUrl(t *testing.T) {
	type fields struct {
		Context     echo.Context
		profileData profile.ConnectionProfile
	}
	type args struct {
		peerUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ethrpc.EthRPC
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := EthernitiContext{
				Context:     tt.fields.Context,
				profileData: tt.fields.profileData,
			}
			got, got1, err := context.RecoverEthClientFromTokenOrPeerUrl(tt.args.peerUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("EthernitiContext.RecoverEthClientFromTokenOrPeerUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EthernitiContext.RecoverEthClientFromTokenOrPeerUrl() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EthernitiContext.RecoverEthClientFromTokenOrPeerUrl() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEthernitiContext_ReadConnectionProfileToken(t *testing.T) {
	type fields struct {
		Context     echo.Context
		profileData profile.ConnectionProfile
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := EthernitiContext{
				Context:     tt.fields.Context,
				profileData: tt.fields.profileData,
			}
			if got := context.ReadConnectionProfileToken(); got != tt.want {
				t.Errorf("EthernitiContext.ReadConnectionProfileToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthernitiContext_JSON(t *testing.T) {
	type fields struct {
		Context     echo.Context
		profileData profile.ConnectionProfile
	}
	type args struct {
		code int
		i    interface{}
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
			context := EthernitiContext{
				Context:     tt.fields.Context,
				profileData: tt.fields.profileData,
			}
			if err := context.JSON(tt.args.code, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("EthernitiContext.JSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEthernitiContext_writeContentType(t *testing.T) {
	type fields struct {
		Context     echo.Context
		profileData profile.ConnectionProfile
	}
	type args struct {
		value string
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
			context := &EthernitiContext{
				Context:     tt.fields.Context,
				profileData: tt.fields.profileData,
			}
			context.writeContentType(tt.args.value)
		})
	}
}

func TestEthernitiContext_Blob(t *testing.T) {
	type fields struct {
		Context     echo.Context
		profileData profile.ConnectionProfile
	}
	type args struct {
		code        int
		contentType string
		b           []byte
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
			context := &EthernitiContext{
				Context:     tt.fields.Context,
				profileData: tt.fields.profileData,
			}
			if err := context.Blob(tt.args.code, tt.args.contentType, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("EthernitiContext.Blob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEthernitiContext_HTMLBlob(t *testing.T) {
	type fields struct {
		Context     echo.Context
		profileData profile.ConnectionProfile
	}
	type args struct {
		code int
		b    []byte
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
			context := &EthernitiContext{
				Context:     tt.fields.Context,
				profileData: tt.fields.profileData,
			}
			if err := context.HTMLBlob(tt.args.code, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("EthernitiContext.HTMLBlob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
