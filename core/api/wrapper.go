// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"bytes"
	"sync"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/trycatch"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	errorPool   sync.Pool
	successPool sync.Pool
	bufferPool  sync.Pool
)

func init() {
	logger.Debug("loading api wrapper data")
	errorPool = sync.Pool{
		New: func() interface{} {
			return protocol.NewApiError(protocol.StatusBadRequest, []byte{})
		},
	}
	successPool = sync.Pool{
		New: func() interface{} {
			return protocol.NewApiResponse([]byte{}, nil)
		},
	}
	bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// return success response to client context
func SendSuccess(c *echo.Context, logMsg []byte, response interface{}) error {
	logger.Debug("sending success message to client")
	logger.Info(str.UnsafeString(logMsg), response)
	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		ToSuccess(logMsg, response),
	)
}

func SendSuccessPool(c *echo.Context, logMsg []byte, response interface{}) error {
	logger.Debug("sending success message to client")
	logger.Info(str.UnsafeString(logMsg), response)
	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		ToSuccessPool(logMsg, response),
	)
}

// return success blob response to client context
func SendSuccessBlob(c *echo.Context, raw []byte) error {
	logger.Debug("sending success blob message to client")
	logger.Info(str.UnsafeString(raw))
	return c.FastBlob(protocol.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, raw)
}

func Success(c *echo.Context, msg []byte, result []byte) error {
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
	return c.FastBlob(protocol.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func ToSuccessPool(msg []byte, result interface{}) []byte {
	logger.Debug("converting data to success message")
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
	logger.Debug("converting data to success payload")
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
	logger.Debug("converting api to error payload")
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
	logger.Debug("converting data to error payload")
	var item protocol.ApiError
	item.Message = msg
	item.Code = code
	b := bufferPool.Get().(*bytes.Buffer)
	rawBytes := item.Bytes(b)
	// put item back to the pool
	bufferPool.Put(b)
	return rawBytes
}

func ErrorStr(c *echo.Context, msg []byte) error {
	logger.Debug("converting error string to payload")
	logger.Error(str.UnsafeString(msg))
	rawBytes := toErrorPool(msg)
	return c.FastBlob(protocol.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func Error(c *echo.Context, err error) error {
	logger.Debug("converting error to payload")
	return ErrorStr(c, str.UnsafeBytes(err.Error()))
}

func ErrorCode(c *echo.Context, code int, err error) error {
	logger.Debug("converting error with code to error payload")
	logger.Error(err)
	rawBytes := toError(code, str.UnsafeBytes(err.Error()))
	return c.FastBlob(code, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func StackError(c *echo.Context, stackErr trycatch.Error) error {
	logger.Debug("converting stack error to error payload")
	logger.Error(stackErr)
	rawBytes := toError(protocol.StatusBadRequest, str.UnsafeBytes(stackErr.Error()))
	return c.FastBlob(protocol.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}
