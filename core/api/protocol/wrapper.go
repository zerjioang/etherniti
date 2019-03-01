package protocol

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/trycatch"
	"github.com/zerjioang/etherniti/core/util"
)

func Success(c echo.Context, msg string, result string) error {
	rawBytes := util.GetJsonBytes(NewApiResponse(msg, result))
	return c.JSONBlob(http.StatusOK, rawBytes)
}

func ErrorStr(c echo.Context, str string) error {
	logger.Error(str)
	rawBytes := util.GetJsonBytes(NewApiError(http.StatusBadRequest, str))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func Error(c echo.Context, err error) error {
	logger.Error(err)
	rawBytes := util.GetJsonBytes(NewApiError(http.StatusBadRequest, err.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func StackError(c echo.Context, stackErr trycatch.Error) error {
	logger.Error(stackErr)
	rawBytes := util.GetJsonBytes(NewApiError(http.StatusBadRequest, stackErr.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}
