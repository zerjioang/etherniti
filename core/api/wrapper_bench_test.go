// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"errors"
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/lib/stack"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"
)

var (
	testSerializer, _ = encoding.EncodingModeSelector(protocol.ModeJson)
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
		//disable logging
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		c := shared.AdquireContext(common.NewContext(echo.New()))
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccess(c, msg, "")
		}
		shared.ReleaseContext(c)
	})
	b.Run("send-success-pool", func(b *testing.B) {
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		c := shared.AdquireContext(common.NewContext(echo.New()))
		//disable logging
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccessPool(c, msg, "")
		}
		shared.ReleaseContext(c)
	})
	b.Run("send-success-blob", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		c := shared.AdquireContext(common.NewContext(echo.New()))
		data := []byte(`{}`)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccessBlob(c, data)
		}
		shared.ReleaseContext(c)
	})
	b.Run("success", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		c := shared.AdquireContext(common.NewContext(echo.New()))
		data := []byte{}
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Success(c, msg, data)
		}
		shared.ReleaseContext(c)
	})
	b.Run("error-str", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		msg := []byte("this is an standard error message working as example")
		c := shared.AdquireContext(common.NewContext(echo.New()))
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ErrorBytes(c, msg)
		}
		shared.ReleaseContext(c)
	})
	b.Run("error", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		c := shared.AdquireContext(common.NewContext(echo.New()))
		e := errors.New("no error")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Error(c, e)
		}
		shared.ReleaseContext(c)
	})
	b.Run("error-code", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		c := shared.AdquireContext(common.NewContext(echo.New()))
		e := errors.New("no error")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ErrorCode(c, 400, e)
		}
		shared.ReleaseContext(c)
	})
	b.Run("stack-error", func(b *testing.B) {
		//disable logging
		logger.Enabled(false)
		c := shared.AdquireContext(common.NewContext(echo.New()))
		e := stack.New("no error")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = StackError(c, e)
		}
		shared.ReleaseContext(c)
	})
}
