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
	logger.ErrorLog.Error(str)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, str))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func Error(c echo.Context, err error) error {
	logger.ErrorLog.Error(err)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, err.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func StackError(c echo.Context, stackErr trycatch.Error) error {
	logger.ErrorLog.Error(stackErr)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, stackErr.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func GetClient(context *server.EthernitiContext) ethrpc.EthRPC {
	client := ethrpc.NewDefaultRPC("http://127.0.0.1:8545")
	return client
}
