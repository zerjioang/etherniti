// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContractNameSpaceController(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		cns := NewContractNameSpaceController()
		assert.NotNil(t, cns)
	})
}
