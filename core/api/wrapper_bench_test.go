package api

import (
	"errors"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/trycatch"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"testing"
)
func BenchmarkWrapper(b *testing.B) {
	b.Run("send-success", func(b *testing.B) {
		ctx := common.NewContext(echo.New())
		//disable logging
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccess(ctx, "message", "")
		}
	})
	b.Run("send-success-blob", func(b *testing.B) {
		ctx := common.NewContext(echo.New())
		//disable logging
		logger.Enabled(false)
		data := []byte(`{}`)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SendSuccessBlob(ctx, data)
		}
	})
	b.Run("success", func(b *testing.B) {
		ctx := common.NewContext(echo.New())
		//disable logging
		logger.Enabled(false)
		data := `{}`
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Success(ctx, "result", data)
		}
	})
	b.Run("error-str", func(b *testing.B) {
		ctx := common.NewContext(echo.New())
		//disable logging
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ErrorStr(ctx, "result")
		}
	})
	b.Run("error", func(b *testing.B) {
		ctx := common.NewContext(echo.New())
		e := errors.New("no error")
		//disable logging
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Error(ctx, e)
		}
	})
	b.Run("error-code", func(b *testing.B) {
		ctx := common.NewContext(echo.New())
		e := errors.New("no error")
		//disable logging
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ErrorCode(ctx, 400, e)
		}
	})
	b.Run("stack-error", func(b *testing.B) {
		ctx := common.NewContext(echo.New())
		e := trycatch.New("no error")
		//disable logging
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = StackError(ctx, e)
		}
	})
}