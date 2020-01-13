// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"errors"
	"sync"
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

func TestWrapper(t *testing.T) {
	t.Run("send-error", func(t *testing.T) {
		c := shared.AdquireContext(common.NewContext(echo.New()))
		err := Error(c, errors.New("test-error"))
		assert.Nil(t, err)
		shared.ReleaseContext(c)
	})
	t.Run("send-error-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				c := shared.AdquireContext(common.NewContext(echo.New()))
				err := Error(c, errors.New("test-error"))
				assert.Nil(t, err)
				g.Done()
				shared.ReleaseContext(c)
			}()
		}
		g.Wait()
	})
	t.Run("send-success", func(t *testing.T) {
		c := shared.AdquireContext(common.NewContext(echo.New()))
		err := SendSuccess(c, []byte("message"), "")
		assert.Nil(t, err)
		shared.ReleaseContext(c)
	})
	t.Run("send-success-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				c := shared.AdquireContext(common.NewContext(echo.New()))
				err := SendSuccess(c, []byte("message"), "")
				assert.Nil(t, err)
				g.Done()
				shared.ReleaseContext(c)
			}()
		}
		g.Wait()
	})
	t.Run("send-success-blob", func(t *testing.T) {
		c := shared.AdquireContext(common.NewContext(echo.New()))
		err := SendSuccessBlob(c, nil)
		assert.Nil(t, err)
		shared.ReleaseContext(c)
	})
	t.Run("send-success-blob-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				c := shared.AdquireContext(common.NewContext(echo.New()))
				err := SendSuccessBlob(c, nil)
				assert.Nil(t, err)
				g.Done()
				shared.ReleaseContext(c)
			}()
		}
		g.Wait()
	})
}
