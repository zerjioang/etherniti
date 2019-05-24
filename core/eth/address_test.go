// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zerjioang/etherniti/core/eth/rpc"
)

const (
	ganacheUITestEndpoint   = "HTTP://127.0.0.1:9545"
	ganacheUIAddress0       = "0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102"
	ganacheCliTestEndpoint  = "HTTP://127.0.0.1:9545"
	ganacheCLIAddress0      = "0xa156Cf364ff355c5727AC79e5363377b6F726129"
	ganacheCLIAddressoEIP55 = "0xa156Cf364Ff355c5727aC79E5363377b6f726129"
	// run tests using ganache cli
	address0            = ganacheCLIAddress0
	ganacheTestEndpoint = ganacheCliTestEndpoint
)

func TestConvertAddress(t *testing.T) {
	addr := ConvertAddress(address0)
	t.Log("address converted", addr.Hex())
	assert.Equal(t, addr.Hex(), ganacheCLIAddressoEIP55, "failed to convert account")
}

func TestIsValidAddress(t *testing.T) {
	result := IsValidAddress(address0)
	assert.True(t, result)
}

func TestGetAccountBalance(t *testing.T) {
	// define the client
	cli := ethrpc.NewDefaultRPC(ganacheTestEndpoint, false)
	expected := big.NewInt(0)
	expected, _ = expected.SetString("100000000000000000000", 10)
	balance, err := cli.EthGetBalance(address0, "latest")
	if err != nil {
		t.Error("failed to get the client", err)
	} else {
		t.Log("readed account balance", balance)
		if balance.Cmp(expected) != 0 {
			t.Error("failed to get balance for ganache account[0]")
		}
	}
}

func TestGetAccountBalanceAtBlock(t *testing.T) {
	// define the client
	cli := ethrpc.NewDefaultRPC(ganacheTestEndpoint, true)
	expected := big.NewInt(0)
	expected, _ = expected.SetString("100000000000000000000", 10)
	balance, err := cli.EthGetBalance(address0, "0")
	if err != nil {
		t.Error("failed to get the client", err)
	} else {
		t.Log("readed account balance", balance)
		if balance.Cmp(expected) != 0 {
			t.Error("failed to get balance for ganache account[0]")
		}
	}
}
