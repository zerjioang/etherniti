// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package token

import (
	"math/big"
	"reflect"
	"testing"

	ethrpc "github.com/zerjioang/etherniti/core/eth/rpc"
)

func TestNewToken(t *testing.T) {
	tests := []struct {
		name string
		want ERC20Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC20Token_Address(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	type args struct {
		address string
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
			t := &ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			t.Address(tt.args.address)
		})
	}
}

func TestERC20Token_Abi(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	type args struct {
		abi string
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
			t := &ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			t.Abi(tt.args.abi)
		})
	}
}

func TestERC20Token_Client(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	type args struct {
		cli ethrpc.EthRPC
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
			t := &ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			t.Client(tt.args.cli)
		})
	}
}

func TestERC20Token_getParams(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			if got := t.getParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ERC20Token.getParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC20Token_Name(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			got, err := t.Name()
			if (err != nil) != tt.wantErr {
				t.Errorf("ERC20Token.Name() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ERC20Token.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC20Token_Symbol(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			got, err := t.Symbol()
			if (err != nil) != tt.wantErr {
				t.Errorf("ERC20Token.Symbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ERC20Token.Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC20Token_Decimals(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint8
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			got, err := t.Decimals()
			if (err != nil) != tt.wantErr {
				t.Errorf("ERC20Token.Decimals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ERC20Token.Decimals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC20Token_TotalSupply(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	tests := []struct {
		name    string
		fields  fields
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			got, err := t.TotalSupply()
			if (err != nil) != tt.wantErr {
				t.Errorf("ERC20Token.TotalSupply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ERC20Token.TotalSupply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC20Token_Call(t *testing.T) {
	type fields struct {
		address     string
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
		abi         string
		cli         ethrpc.EthRPC
	}
	type args struct {
		method string
		params []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := ERC20Token{
				address:     tt.fields.address,
				name:        tt.fields.name,
				symbol:      tt.fields.symbol,
				decimals:    tt.fields.decimals,
				totalSupply: tt.fields.totalSupply,
				abi:         tt.fields.abi,
				cli:         tt.fields.cli,
			}
			got, err := t.Call(tt.args.method, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ERC20Token.Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ERC20Token.Call() = %v, want %v", got, tt.want)
			}
		})
	}
}
