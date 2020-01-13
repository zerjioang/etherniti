// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package wallet

import (
	"testing"

	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/go-hpc/lib/eth"
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
		var response dto.AccountResponse
		response.Address = address.Hex()
		response.Key = privateKey
		assert.NotNil(t, response)
		assert.NotNil(t, response.Address)
		assert.NotNil(t, response.Key)
	})
}
