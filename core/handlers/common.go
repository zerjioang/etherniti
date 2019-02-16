package handlers

import (
	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/util"
	"net/http"
)

func ErrorStr(c echo.Context, str string) error {
	rawBytes := util.GetJsonBytes( api.NewApiError(http.StatusBadRequest, str) )
	return c.JSONBlob(http.StatusOK, rawBytes)
}

func Error(c echo.Context, err error) error {
	rawBytes := util.GetJsonBytes( api.NewApiError(http.StatusBadRequest, err.Error()) )
	return c.JSONBlob(http.StatusOK, rawBytes)
}