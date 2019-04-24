// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"net/http"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/trycatch"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// return success response to client context
func SendSuccess(c echo.ContextInterface, logMsg string, response interface{}) error {
	logger.Debug("sending success message to client")
	logger.Info(logMsg)
	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		ToSuccess(logMsg, response),
	)
}

// return success blob response to client context
func SendSuccessBlob(c echo.ContextInterface, raw []byte) error {
	logger.Debug("sending success BLOB message to client")
	return c.FastBlob(protocol.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, raw)
}

func Success(c echo.ContextInterface, msg string, result string) error {
	logger.Debug("sending success message to client")
	rawBytes := str.GetJsonBytes(protocol.NewApiResponse(msg, result))
	return c.FastBlob(http.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func ToSuccess(msg string, result interface{}) []byte {
	logger.Debug("generating success byte array")
	rawBytes := str.GetJsonBytes(protocol.NewApiResponse(msg, result))
	return rawBytes
}

func ErrorStr(c echo.ContextInterface, msg string) error {
	logger.Error(msg)
	rawBytes := str.GetJsonBytes(protocol.NewApiError(http.StatusBadRequest, msg))
	return c.FastBlob(http.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func Error(c echo.ContextInterface, err error) error {
	logger.Error(err)
	apierr := protocol.NewApiError(http.StatusBadRequest, err.Error())
	rawBytes := str.GetJsonBytes(apierr)
	return c.FastBlob(http.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func ErrorCode(c echo.ContextInterface, code int, err error) error {
	logger.Error(err)
	apierr := protocol.NewApiError(code, err.Error())
	rawBytes := str.GetJsonBytes(apierr)
	return c.FastBlob(code, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func StackError(c echo.ContextInterface, stackErr trycatch.Error) error {
	logger.Error(stackErr)
	rawBytes := str.GetJsonBytes(protocol.NewApiError(http.StatusBadRequest, stackErr.Error()))
	return c.FastBlob(http.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}
