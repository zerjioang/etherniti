package e2e_test

import (
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/core/listener/middleware"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func BenchmarkAPI(b *testing.B) {
	b.Run("sha3-local", func(b *testing.B) {
		// Setup
		e := common.NewServer(middleware.ConfigureServerRoutes)
		logger.Enabled(false)
		postData := `{"data": "hello-world"}`

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		//start benchmarking
		for n := 0; n < b.N; n++ {
			req := httptest.NewRequest(http.MethodPost, constants.ApiVersion+"/web3/ganache/sha3/local", strings.NewReader(postData))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
		}
	})
}
