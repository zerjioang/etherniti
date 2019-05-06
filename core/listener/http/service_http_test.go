// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/http"
)

func TestHttpListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		http.NewHttpListener()
	})
	t.Run("run", func(t *testing.T) {
		s := http.NewHttpListener()
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
	})
	t.Run("request-status", func(t *testing.T) {
		s := http.NewHttpListener()
		// run the socket server
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
	})
}
