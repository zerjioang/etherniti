// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package socket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnixSocketListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		NewSocketListener()
	})
	t.Run("run", func(t *testing.T) {
		s := NewSocketListener()
		s.RunMode("/tmp/go.sock", true)
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
	})
	t.Run("request-status", func(t *testing.T) {
		s := NewSocketListener()
		// run the socket server
		s.RunMode("/tmp/go.sock", true)
		notifier := make(chan error, 1)
		s.Listen(notifier)
		err := <-notifier
		assert.Nil(t, err)
		// send GET style request for v1/public for welcome message
		cli := socketHttpClient("/tmp/go.sock")
		resp, err := cli.Get("http://unix/v1/public")
		t.Log("response", resp)
		t.Log("error", err)
	})
}
