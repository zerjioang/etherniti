// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/etherniti/core/modules/cache"
)

func TestNewNetworkController(t *testing.T) {
	tests := []struct {
		name string
		want NetworkController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNetworkController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNetworkController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetworkController_SetPeer(t *testing.T) {
	type fields struct {
		cache       *cache.MemoryCache
		peer        string
		networkName string
	}
	type args struct {
		peerLocation string
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
			ctl := &NetworkController{
				cache:       tt.fields.cache,
				peer:        tt.fields.peer,
				networkName: tt.fields.networkName,
			}
			ctl.SetPeer(tt.args.peerLocation)
		})
	}
}

func TestNetworkController_SetTargetName(t *testing.T) {
	type fields struct {
		cache       *cache.MemoryCache
		peer        string
		networkName string
	}
	type args struct {
		networkName string
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
			ctl := &NetworkController{
				cache:       tt.fields.cache,
				peer:        tt.fields.peer,
				networkName: tt.fields.networkName,
			}
			ctl.SetTargetName(tt.args.networkName)
		})
	}
}

func TestNetworkController_RegisterRouters(t *testing.T) {
	type fields struct {
		cache       *cache.MemoryCache
		peer        string
		networkName string
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
			ctl := NetworkController{
				cache:       tt.fields.cache,
				peer:        tt.fields.peer,
				networkName: tt.fields.networkName,
			}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
