// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package https_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zerjioang/etherniti/core/listener/https"
)

func TestHttpListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		l := https.NewHttpsListener()
		assert.NotNil(t, l)
	})
	t.Run("instantiation-custom", func(t *testing.T) {
		l := https.NewHttpsListenerCustom()
		assert.NotNil(t, l)
	})
}
