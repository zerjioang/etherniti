// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"net/http"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/trycatch"
)

// return success response to client context
func SendSuccess(c echo.Context, logMsg string, response interface{}) error {
	logger.Debug("sending success message to client")
	logger.Info(logMsg)
	raw := ToSuccess(logMsg, response)
	return c.JSONBlob(protocol.StatusOK, raw)
}

// return success blob response to client context
func SendSuccessBlob(c echo.Context, raw []byte) error {
	logger.Debug("sending success BLOB message to client")
	return c.JSONBlob(protocol.StatusOK, raw)
}

func Success(c echo.Context, msg string, result string) error {
	logger.Debug("sending success message to client")
	rawBytes := str.GetJsonBytes(protocol.NewApiResponse(msg, result))
	return c.JSONBlob(http.StatusOK, rawBytes)
}

func ToSuccess(msg string, result interface{}) []byte {
	logger.Debug("generating success byte array")
	rawBytes := str.GetJsonBytes(protocol.NewApiResponse(msg, result))
	return rawBytes
}

func ErrorStr(c echo.Context, msg string) error {
	logger.Error(msg)
	rawBytes := str.GetJsonBytes(protocol.NewApiError(http.StatusBadRequest, msg))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func Error(c echo.Context, err error) error {
	logger.Error(err)
	apierr := protocol.NewApiError(http.StatusBadRequest, err.Error())
	rawBytes := str.GetJsonBytes(apierr)
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}

func StackError(c echo.Context, stackErr trycatch.Error) error {
	logger.Error(stackErr)
	rawBytes := str.GetJsonBytes(protocol.NewApiError(http.StatusBadRequest, stackErr.Error()))
	return c.JSONBlob(http.StatusBadRequest, rawBytes)
}
