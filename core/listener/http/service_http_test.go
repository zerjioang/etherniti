// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http

import (
	"testing"
	"time"
)

func TestHttpListener(t *testing.T) {
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
	})
}
