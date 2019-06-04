// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ws_test

import (
	"testing"

	ws "github.com/zerjioang/etherniti/core/listener/websocket"
)

func TestWebsocketListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		ws.NewWebsocketListener()
	})
}
