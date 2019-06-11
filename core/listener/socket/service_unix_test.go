// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package socket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnixSocketListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		l := NewSocketListener()
		assert.NotNil(t, l)
	})
	t.Run("run", func(t *testing.T) {
		s := NewSocketListener()
		assert.NotNil(t, s)
		s.RunMode("/tmp/go.sock", true)
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
	})
	t.Run("request-status", func(t *testing.T) {
		s := NewSocketListener()
		assert.NotNil(t, s)
		// run the socket server
		s.RunMode("/tmp/go.sock", true)
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
		// send GET style request for v1/hi for welcome message
		cli := socketHttpClient("/tmp/go.sock")
		resp, err := cli.Get("http://unix/v1/hi")
		t.Log("response", resp)
		t.Log("error", err)
	})
}
