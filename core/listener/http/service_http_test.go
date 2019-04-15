// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http_test

import (
	"github.com/zerjioang/etherniti/core/listener/http"
	"testing"
	"time"
)

func TestHttpListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		http.NewHttpListener()
	})
	t.Run("run", func(t *testing.T) {
		s := http.NewHttpListener()
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		time.Sleep(200000 * time.Second)
	})
	t.Run("request-status", func(t *testing.T) {
		s := http.NewHttpListener()
		// run the socket server
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		// wait one second to bootup
		time.Sleep(1 * time.Second)
	})
}
