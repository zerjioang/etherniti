// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/trycatch"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	errorPool   *sync.Pool
	successPool *sync.Pool
	bufferPool  *sync.Pool
)

func init() {
	errorPool = &sync.Pool{
		New: func() interface{} {
			return protocol.NewApiError(http.StatusBadRequest, []byte{})
		},
	}
	successPool = &sync.Pool{
		New: func() interface{} {
			return protocol.NewApiResponse([]byte{}, nil)
		},
	}
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// return success response to client context
func SendSuccess(c echo.ContextInterface, logMsg []byte, response interface{}) error {
	logger.Debug("sending success message to client")
	logger.Info(str.UnsafeString(logMsg), response)
	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		ToSuccess(logMsg, response),
	)
}

func SendSuccessPool(c echo.ContextInterface, logMsg []byte, response interface{}) error {
	logger.Debug("sending success message to client")
	logger.Info(str.UnsafeString(logMsg), response)
	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		ToSuccessPool(logMsg, response),
	)
}

// return success blob response to client context
func SendSuccessBlob(c echo.ContextInterface, raw []byte) error {
	logger.Debug("sending success blob message to client")
	logger.Info( str.UnsafeString(raw) )
	return c.FastBlob(protocol.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, raw)
}

func Success(c echo.ContextInterface, msg []byte, result []byte) error {
	logger.Debug("sending success message to client")
	logger.Debug(str.UnsafeString(msg), str.UnsafeString(result))
	//get item from pool
	item := successPool.Get().(*protocol.ApiResponse)
	item.Message = msg
	item.Result = result
	b := bufferPool.Get().(*bytes.Buffer)
	rawBytes := item.Bytes(b)
	// put item back to the pool
	bufferPool.Put(b)
	successPool.Put(item)
	return c.FastBlob(http.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func ToSuccessPool(msg []byte, result interface{}) []byte {
	//get item from pool
	item := successPool.Get().(*protocol.ApiResponse)
	item.Message = msg
	item.Result = result
	b := bufferPool.Get().(*bytes.Buffer)
	rawBytes := item.Bytes(b)
	// put item back to the pool
	bufferPool.Put(b)
	successPool.Put(item)

	return rawBytes
}

func ToSuccess(msg []byte, result interface{}) []byte {
	//get item from pool
	var item protocol.ApiResponse
	item.Code = 200
	item.Message = msg
	item.Result = result
	b := bufferPool.Get().(*bytes.Buffer)
	rawBytes := item.Bytes(b)
	// put item back to the pool
	bufferPool.Put(b)
	return rawBytes
}

func toErrorPool(msg []byte) []byte {
	//get item from pool
	item := errorPool.Get().(*protocol.ApiError)
	item.Message = msg
	b := bufferPool.Get().(*bytes.Buffer)
	rawBytes := item.Bytes(b)
	// put item back to the pool
	bufferPool.Put(b)
	errorPool.Put(item)
	return rawBytes
}

func toError(code int, msg []byte) []byte {
	var item protocol.ApiError
	item.Message = msg
	item.Code = code
	b := bufferPool.Get().(*bytes.Buffer)
	rawBytes := item.Bytes(b)
	// put item back to the pool
	bufferPool.Put(b)
	return rawBytes
}

func ErrorStr(c echo.ContextInterface, msg []byte) error {
	logger.Error(str.UnsafeString(msg))
	rawBytes := toErrorPool(msg)
	return c.FastBlob(http.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func Error(c echo.ContextInterface, err error) error {
	return ErrorStr(c, str.UnsafeBytes(err.Error()))
}

func ErrorCode(c echo.ContextInterface, code int, err error) error {
	logger.Error(err)
	rawBytes := toError(code, str.UnsafeBytes(err.Error()))
	return c.FastBlob(code, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func StackError(c echo.ContextInterface, stackErr trycatch.Error) error {
	logger.Error(stackErr)
	rawBytes := toError(http.StatusBadRequest, str.UnsafeBytes(stackErr.Error()))
	return c.FastBlob(http.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}
