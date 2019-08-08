// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/shared/protocol"
	"testing"
)

func TestWalletController(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {

	})
	t.Run("generate-key", func(t *testing.T) {
		// Create an account
		private, err := eth.GenerateNewKey()
		assert.Nil(t, err)
		address := eth.GetAddressFromPrivateKey(private)
		privateKey := eth.GetPrivateKeyAsEthString(private)
		var response protocol.AccountResponse
		response.Address = address.Hex()
		response.Key = privateKey
	})
}