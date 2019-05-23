// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package https_test

import (
	"testing"

	"github.com/zerjioang/etherniti/core/listener/https"
)

func TestHttpListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		https.NewHttpsListener()
	})
}
