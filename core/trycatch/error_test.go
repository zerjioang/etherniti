// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package trycatch

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	t.Run("instantiate-nil", func(t *testing.T) {
		er := Nil()
		assert.True(t, er.None())
	})
	t.Run("instantiate-not-nil", func(t *testing.T) {
		er := New("test")
		assert.True(t, er.Occur())
	})
	t.Run("instantiate-from-err", func(t *testing.T) {
		er := errors.New("default")
		stackErr := New(er.Error())
		assert.NotNil(t, stackErr)
		assert.True(t, stackErr.Occur())
	})
}
