// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"math/big"
	"testing"
)

const (
	ganacheUITestEndpoint  = "HTTP://127.0.0.1:9545"
	ganacheUIAddress0      = "0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102"
	ganacheCliTestEndpoint = "HTTP://127.0.0.1:9545"
	ganacheCLIAddress0     = "0xa156Cf364ff355c5727AC79e5363377b6F726129"
	// run tests using ganache cli
	address0            = ganacheCLIAddress0
	ganacheTestEndpoint = ganacheCliTestEndpoint
)

func TestConvertAddress(t *testing.T) {
	addr := ConvertAddress(address0)
	t.Log("address converted", addr.Hex())
	if addr.Hex() != address0 {
		t.Error("failed to convert account")
	}
}

func TestGetAccountBalance(t *testing.T) {
	// define the address
	addr := ConvertAddress(address0)
	// define the client
	ganacheClient, err := GetEthereumClient(HttpClient, ganacheTestEndpoint)
	if err != nil {
		t.Error("failed to get the client", err)
	} else if ganacheClient == nil {
		t.Error("failed to get a valid client")
	}
	balance, bErr := GetAccountBalance(ganacheClient, addr)
	if bErr != nil {
		t.Error("failed to get account balance", bErr)
	} else {
		expected := big.NewInt(0)
		expected, _ = expected.SetString("100000000000000000000", 10)
		t.Log("readed account balance", balance)
		if balance.Cmp(expected) != 0 {
			t.Error("failed to get balance for ganache account[0]")
		}
	}
}

func TestGetAccountBalanceAtBlock(t *testing.T) {
	// define the address
	addr := ConvertAddress(address0)
	// define the client
	ganacheClient, err := GetEthereumClient(HttpClient, ganacheTestEndpoint)
	if err != nil {
		t.Error("failed to get the client", err)
	} else if ganacheClient == nil {
		t.Error("failed to get a valid client")
	}
	balance, bErr := GetAccountBalanceAtBlock(ganacheClient, addr, big.NewInt(0))
	if bErr != nil {
		t.Error("failed to get account balance", err)
	} else {
		expected := big.NewInt(0)
		expected, _ = expected.SetString("100000000000000000000", 10)
		t.Log(balance)
		if balance.Cmp(expected) != 0 {
			t.Error("failed to get balance for ganache account[0]")
		}
	}
}

func TestToEth(t *testing.T) {
	// define the address
	addr := ConvertAddress(address0)
	// define the client
	ganacheClient, err := GetEthereumClient(HttpClient, ganacheTestEndpoint)
	if err != nil {
		t.Error("failed to get the client", err)
	} else if ganacheClient == nil {
		t.Error("failed to get a valid client")
	}
	balance, bErr := GetAccountBalance(ganacheClient, addr)
	if bErr != nil {
		t.Error("failed to get account balance", err)
	} else {
		ethValue := ToEth(*balance)
		t.Log("ETH value", ethValue)
	}
}
