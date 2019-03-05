// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http

import (
	"testing"
	"time"
)

func TestUnixSocketListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		NewHttpListener()
	})
	t.Run("run", func(t *testing.T) {
		s := NewHttpListener()
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		time.Sleep(200000 * time.Second)
	})
	t.Run("request-status", func(t *testing.T) {
		s := NewHttpListener()
		// run the socket servre
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		// wait one second to bootup
		time.Sleep(1 * time.Second)
		// send GET style request for v1/public for welcome message
		cli := socketHttpClient("/tmp/go.sock")
		resp, err := cli.Get("http://unix/v1/public")
		t.Log("response", resp)
		t.Log("error", err)
	})
}
