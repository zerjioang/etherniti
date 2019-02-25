package handlers

import (
	"net/http"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"
	"github.com/zerjioang/etherniti/core/trycatch"

	"github.com/zerjioang/etherniti/core/eth/rpc"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/util"
)

func Success(c echo.Context, msg string, result string) error {
	rawBytes := util.GetJsonBytes(api.NewApiResponse(msg, result))
	return c.JSONBlob(http.StatusOK, rawBytes)
}

func ErrorStr(c echo.Context, str string) error {
	logger.Error(str)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, str))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func Error(c echo.Context, err error) error {
	logger.Error(err)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, err.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func StackError(c echo.Context, stackErr trycatch.Error) error {
	logger.Error(stackErr)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, stackErr.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func GetClient(context *server.EthernitiContext) ethrpc.EthRPC {
	client := ethrpc.NewDefaultRPC("http://127.0.0.1:8545")
	return client
}

// cache middleware function.
/*
The Cache-Control header is the most important header
to set as it effectively ‘switches on’ caching in
the browser. With this header in place, and set with
a value that enables caching, the browser will cache
the file for as long as specified. Without this header
the browser will re-request the file on each
subsequent request.

public resources can be cached not only by the
end-user’s browser but also by any intermediate
proxies that may be serving many other users as well.

private resources are bypassed by intermediate
proxies and can only be cached by the end-client.

The max-age value sets a timespan for how
long to cache the resource (in seconds).
 */
func Cached(c echo.Context, cacheValid bool) (int, echo.Context) {
	// add cache headers
	r := c.Response()
	h := r.Header()
	h.Set("Cache-Control", "public, max-age=86400") // 24h cache
	//h.Set("Cache-Control", "private")
	if cacheValid {
		//cached item is still valid, so return a not modified
		r.Status = http.StatusOK // http.StatusNotModified
	} else {
		// cached data set as invalid, return 200 ok in order to update
		r.Status = http.StatusOK
	}
	return r.Status, c
}