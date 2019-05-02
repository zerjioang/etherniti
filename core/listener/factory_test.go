// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package listener

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/shared/def/listener"
)

func TestFactoryListener(t *testing.T) {
	t.Run("factory-http", func(t *testing.T) {
		ln := FactoryListener(listener.HttpMode)
		assert.NotNil(t, ln)
	})
	t.Run("factory-https", func(t *testing.T) {
		ln := FactoryListener(listener.HttpsMode)
		assert.NotNil(t, ln)
	})
	t.Run("factory-unix", func(t *testing.T) {
		ln := FactoryListener(listener.UnixMode)
		assert.NotNil(t, ln)
	})
	t.Run("factory-other", func(t *testing.T) {
		ln := FactoryListener(listener.UnknownMode)
		assert.NotNil(t, ln)
	})
}
