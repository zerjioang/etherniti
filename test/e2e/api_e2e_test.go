package e2e_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/core/listener/middleware"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestAPI(t *testing.T) {
	t.Run("sha3-local", func(t *testing.T) {
		// Setup
		e := common.NewServer(middleware.ConfigureServerRoutes)
		postData := `{"data": "hello-world"}`

		req := httptest.NewRequest(http.MethodPost, constants.ApiVersion+"/web3/ganache/sha3/local", strings.NewReader(postData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		//execute the request
		e.ServeHTTP(rec, req)
		//read the response
		raw, err := ioutil.ReadAll(rec.Body)
		assert.NoError(t, err)
		assert.NotNil(t, raw)
		assert.Equal(t, string(raw), `{"msg":"sha3","data":"0xd41bad2284cfa351467b5db9418bbe3a5c02162c02ee585f07e5553d823ebad9"}`)
	})
}
