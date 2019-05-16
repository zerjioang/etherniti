// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cns

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/modules/stack"
)

func TestContractNameSystem(t *testing.T) {
	t.Run("e2e-test", func(t *testing.T) {
		cns := NewContractNameSystem()
		assert.NotNil(t, cns)

		contract := ContractInfo{}
		contract.Name = "test"
		contract.Description = "this is a demo contract"
		contract.Address = "0xf17f52151EbEF6C7334FAD080c5704D77216b732"
		contract.Version = "1.2"

		cns.Register(contract)

		response, success := cns.Resolve("test-1.2")
		assert.Equal(t, response.Version, "1.2")
		assert.Equal(t, response.Address, "0xf17f52151EbEF6C7334FAD080c5704D77216b732")
		assert.Equal(t, response.Description, "this is a demo contract")
		assert.Equal(t, response.Name, "test")
		assert.Equal(t, success, true)
		t.Log(cns.Resolve("test-02"))
	})
}

func TestContractInfo_Id(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		Address     string
		Version     string
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
			c := ContractInfo{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Address:     tt.fields.Address,
				Version:     tt.fields.Version,
			}
			if got := c.Id(); got != tt.want {
				t.Errorf("ContractInfo.Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContractInfo_Validate(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		Address     string
		Version     string
	}
	tests := []struct {
		name   string
		fields fields
		want   stack.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ContractInfo{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Address:     tt.fields.Address,
				Version:     tt.fields.Version,
			}
			if got := c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContractInfo.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContractNameSystem(t *testing.T) {
	tests := []struct {
		name string
		want ContractNameSystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContractNameSystem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContractNameSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContractNameSystemPtr(t *testing.T) {
	tests := []struct {
		name string
		want *ContractNameSystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContractNameSystemPtr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContractNameSystemPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContractNameSystem_Register(t *testing.T) {
	type fields struct {
		data *sync.Map
	}
	type args struct {
		info ContractInfo
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
			ns := &ContractNameSystem{
				data: tt.fields.data,
			}
			ns.Register(tt.args.info)
		})
	}
}

func TestContractNameSystem_Unregister(t *testing.T) {
	type fields struct {
		data *sync.Map
	}
	type args struct {
		id string
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
			ns := &ContractNameSystem{
				data: tt.fields.data,
			}
			ns.Unregister(tt.args.id)
		})
	}
}

func TestContractNameSystem_Resolve(t *testing.T) {
	type fields struct {
		data *sync.Map
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ContractInfo
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &ContractNameSystem{
				data: tt.fields.data,
			}
			got, got1 := ns.Resolve(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContractNameSystem.Resolve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ContractNameSystem.Resolve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
