package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/controllers"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestIndexHandler(t *testing.T) {
	// create our server
	// Setup
	e := echo.New()

	// Create a request to pass to our handler.
	// We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", constants.ApiVersion+"/hi", nil)
	assert.Nil(t, err)
	assert.NotNil(t, req)
	// set content type to json
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-Real-IP", "127.0.0.1")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rec := httptest.NewRecorder()

	//build this test execution context
	c := e.NewContext(req, rec)
	c.Preload()

	// build our controller
	ctl := controllers.NewIndexController()
	runErr := ctl.Index(c)
	assert.Nil(t, runErr)

	// Check the status code is what we expect.
	assert.Equal(t, rec.Code, http.StatusOK, "handler returned wrong status code")
	assert.Equal(t, rec.Body.String(), controllers.IndexWelcomeJson, "handler returned unexpected body")
}
