// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"math/big"
	"testing"
)

const (
	ganacheTestEndpoint = "HTTP://127.0.0.1:9545"
)

func TestConvertAddress(t *testing.T) {
	addr := ConvertAddress("0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102")
	if addr.Hex() != "0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102" {
		t.Error("failed to convert account")
	}
}

func TestGetAccountBalance(t *testing.T) {
	// define the address
	addr := ConvertAddress("0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102")
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
		expected := big.NewInt(0)
		expected, _ = expected.SetString("100000000000000000000", 10)
		t.Log(balance)
		if balance.Cmp(expected) != 0 {
			t.Error("failed to get balance for ganache account[0]")
		}
	}
}

func TestGetAccountBalanceAtBlock(t *testing.T) {
	// define the address
	addr := ConvertAddress("0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102")
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
	addr := ConvertAddress("0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102")
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
		ethValue := ToEth(balance)
		t.Log("ETH value", ethValue)
	}
}
