// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestWrapper(t *testing.T) {
	t.Run("send-error", func(t *testing.T) {
		err := Error(common.NewContext(echo.New()), errors.New("test-error"))
		assert.Nil(t, err)
	})
	t.Run("send-error-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				err := Error(common.NewContext(echo.New()), errors.New("test-error"))
				assert.Nil(t, err)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("send-success", func(t *testing.T) {
		err := SendSuccess(common.NewContext(echo.New()), []byte("message"), "")
		assert.Nil(t, err)
	})
	t.Run("send-success-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				err := SendSuccess(common.NewContext(echo.New()), []byte("message"), "")
				assert.Nil(t, err)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("send-success-blob", func(t *testing.T) {
		err := SendSuccessBlob(common.NewContext(echo.New()), nil)
		assert.Nil(t, err)
	})
	t.Run("send-success-blob-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				err := SendSuccessBlob(common.NewContext(echo.New()), nil)
				assert.Nil(t, err)
				g.Done()
			}()
		}
		g.Wait()
	})
}
