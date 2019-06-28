// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mtls_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/mtls"
)

func TestMtlsListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		l := mtls.NewMtlsListener()
		assert.NotNil(t, l)
	})
	t.Run("server-config", func(t *testing.T) {
		l := mtls.NewMtlsListener()
		assert.NotNil(t, l)
		scfg := l.ServerConfig()
		assert.NotNil(t, scfg)
	})
}
