// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cns

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
