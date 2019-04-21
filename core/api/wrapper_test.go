// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"testing"
)

func TestWrapper(t *testing.T) {
	t.Run("send-success", func(t *testing.T) {
		err := SendSuccess(common.NewContext(echo.New()), "message", "")
		assert.Nil(t, err)
	})
	t.Run("send-success-blob", func(t *testing.T) {
		err := SendSuccessBlob(common.NewContext(echo.New()), nil)
		assert.Nil(t, err)
	})
}
