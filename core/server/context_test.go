// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import (
	"testing"
)

func TestNewEthernitiContext(t *testing.T) {
	t.Run("instantiate-nil", func(t *testing.T) {
		_ = NewEthernitiContext(nil)
	})
}