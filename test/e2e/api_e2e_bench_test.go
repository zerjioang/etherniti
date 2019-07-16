package e2e_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/core/listener/middleware"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func BenchmarkAPI(b *testing.B) {
	// BenchmarkAPI/create-new-recorder-4         	100000000	        20.0 ns/op	  50.08 MB/s	       0 B/op	       0 allocs/op
	b.Run("create-new-recorder", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			//execute the request
			_ = httptest.NewRecorder()
		}
	})
	b.Run("sha3-local", func(b *testing.B) {
		// Setup
		e := common.NewServer(middleware.ConfigureServerRoutes)
		postData := `{"data": "hello-world"}`
		// disable logger
		logger.Enabled(false)

		//start benchmarking
		req := httptest.NewRequest(http.MethodPost, constants.ApiVersion+"/web3/ganache/sha3/local", strings.NewReader(postData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			//execute the request
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
		}
	})
}
