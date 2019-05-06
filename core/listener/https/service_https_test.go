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
		https.NewHttpsListener()
	})
	t.Run("run", func(t *testing.T) {
		s := https.NewHttpsListener()
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
	})
	t.Run("request-status", func(t *testing.T) {
		s := https.NewHttpsListener()
		// run the socket server
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
	})
}
