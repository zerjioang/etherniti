// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package memory

import (
	"crypto/ecdsa"
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
	ethrpc "github.com/zerjioang/etherniti/core/eth/rpc"
)

func TestWalletContent_Client(t *testing.T) {
	type fields struct {
		ethAddress       fixtures.Address
		privateKey       ecdsa.PrivateKey
		connectionClient ethrpc.EthRPC
	}
	tests := []struct {
		name   string
		fields fields
		want   ethrpc.EthRPC
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet := WalletContent{
				ethAddress:       tt.fields.ethAddress,
				privateKey:       tt.fields.privateKey,
				connectionClient: tt.fields.connectionClient,
			}
			if got := wallet.Client(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalletContent.Client() = %v, want %v", got, tt.want)
			}
		})
	}
}
