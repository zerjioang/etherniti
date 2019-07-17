// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"errors"
	"testing"

	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto"
	"github.com/zerjioang/etherniti/shared/protocol/io"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/stack"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	testSerializer, _ = ioproto.EncodingModeSelector(io.ModeJson)
)

func BenchmarkWrapper(b *testing.B) {
	b.Run("to-success", func(b *testing.B) {
		data := []byte("message")
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ToSuccess(data, "", testSerializer)
		}
	})

	b.Run("to-error", func(b *testing.B) {
		msg := "this is an standard error message working as example"
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = toError(0, msg, "", testSerializer)
		}
	})
	b.Run("to-error-pool", func(b *testing.B) {
		msg := "this is an standard error message working as example"
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = toErrorPool(msg, testSerializer)
		}
	})
	b.Run("send-success", func(b *testing.B) {
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		ctx := common.NewContext(echo.New())
		//disable logging
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccess(ctx, msg, "")
		}
	})
	b.Run("send-success-pool", func(b *testing.B) {
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		ctx := common.NewContext(echo.New())
		//disable logging
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccessPool(ctx, msg, "")
		}
	})
	b.Run("send-success-blob", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		ctx := common.NewContext(echo.New())
		data := []byte(`{}`)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccessBlob(ctx, data)
		}
	})
	b.Run("success", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		ctx := common.NewContext(echo.New())
		data := []byte{}
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Success(ctx, msg, data)
		}
	})
	b.Run("error-str", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		ctx := common.NewContext(echo.New())
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ErrorBytes(ctx, msg)
		}
	})
	b.Run("error", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		ctx := common.NewContext(echo.New())
		e := errors.New("no error")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Error(ctx, e)
		}
	})
	b.Run("error-code", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		ctx := common.NewContext(echo.New())
		e := errors.New("no error")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ErrorCode(ctx, 400, e)
		}
	})
	b.Run("stack-error", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		ctx := common.NewContext(echo.New())
		e := stack.New("no error")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = StackError(ctx, e)
		}
	})
}
