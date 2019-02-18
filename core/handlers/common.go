package handlers

import (
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/core/server"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/util"
)

func Success(c echo.Context, msg string, result string) error {
	rawBytes := util.GetJsonBytes(api.NewApiResponse(msg, result))
	return c.JSONBlob(http.StatusOK, rawBytes)
}

func ErrorStr(c echo.Context, str string) error {
	log.Error(str)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, str))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func Error(c echo.Context, err error) error {
	log.Error(err)
	rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, err.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

// from incoming http request, it recovers the eth client linked to it
func GetClientInstance(c *server.EthernitiContext) (ethrpc.EthRPC, error) {
	// requestProfileKey := c.Request().Header.Get(config.HttpProfileHeaderkey)
	client := ethrpc.NewDefaultRPC("http://127.0.0.1:8545")
	return client, nil
}

func GetClient(context *server.EthernitiContext) ethrpc.EthRPC {
	client := ethrpc.NewDefaultRPC("http://127.0.0.1:8545")
	return client
}