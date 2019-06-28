// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	t.Run("secure", func(t *testing.T) {
		e := secure(nil)
		assert.NotNil(t, e)
	})
}
