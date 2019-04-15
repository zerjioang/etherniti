// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package https_test

import (
	"github.com/zerjioang/etherniti/core/listener/https"
	"testing"
	"time"
)

func TestHttpListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		https.NewHttpsListener()
	})
	t.Run("run", func(t *testing.T) {
		s := https.NewHttpsListener()
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		time.Sleep(200000 * time.Second)
	})
	t.Run("request-status", func(t *testing.T) {
		s := https.NewHttpsListener()
		// run the socket server
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		// wait one second to bootup
		time.Sleep(1 * time.Second)
	})
}