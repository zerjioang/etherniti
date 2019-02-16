package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/util"
)

func ErrorStr(c echo.Context, str string) error {
	log.Error(str)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, str))
	return c.JSONBlob(http.StatusOK, rawBytes)
}

func Error(c echo.Context, err error) error {
	log.Error(err)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, err.Error()))
	return c.JSONBlob(http.StatusOK, rawBytes)
}
