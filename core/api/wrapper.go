// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"bytes"
	"sync"

	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto"
	"github.com/zerjioang/etherniti/shared/protocol/io"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/stack"
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
	logger.Debug(str.UnsafeString(logMsg), " - ", response)
	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		ToSuccess(logMsg, response, c.ResponseSerializer()),
	)
}

// return success response to client context
func SendRawSuccess(c *echo.Context, content []byte) error {
	logger.Debug("sending success message to client")
	logger.Debug(str.UnsafeString(content))
	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		content,
	)
}

func SendSuccessPool(c *echo.Context, logMsg []byte, v interface{}) error {
	logger.Debug("sending success message to client")
	logger.Info(str.UnsafeString(logMsg), v)

	//generate byte content
	logger.Debug("converting data to success message")
	//get item from pool
	item := successPool.Get().(*protocol.ApiResponse)
	item.Message = str.UnsafeString(logMsg)
	item.Data = v
	rawBytes := ioproto.GetBytesFromSerializer(c.ResponseSerializer(), item)
	// put item back to the pool
	successPool.Put(item)

	return c.FastBlob(
		protocol.StatusOK,
		echo.MIMEApplicationJSONCharsetUTF8,
		rawBytes,
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
	logger.Debug(str.UnsafeString(msg), " , ", str.UnsafeString(result))
	//get item from pool
	item := successPool.Get().(*protocol.ApiResponse)
	item.Message = str.UnsafeString(msg)
	item.Data = result
	rawBytes := ioproto.GetBytesFromSerializer(c.ResponseSerializer(), item)
	// put item back to the pool
	successPool.Put(item)
	return c.FastBlob(protocol.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func ToSuccess(msg []byte, result interface{}, serializer io.Serializer) []byte {
	logger.Debug("converting data to success payload")
	//get item from pool
	var item protocol.ApiResponse
	item.Message = str.UnsafeString(msg)
	item.Data = result
	rawBytes := ioproto.GetBytesFromSerializer(serializer, item)
	// put item back to the pool
	return rawBytes
}

func toErrorPool(msg []byte, serializer io.Serializer) []byte {
	logger.Debug("converting api to error payload")
	//get item from pool
	item := errorPool.Get().(*protocol.ApiError)
	item.Desc = msg
	rawBytes := ioproto.GetBytesFromSerializer(serializer, item)
	// put item back to the pool
	errorPool.Put(item)
	return rawBytes
}

func toError(code int, msg []byte, data []byte, serializer io.Serializer) []byte {
	logger.Debug("converting data to error payload")
	var item protocol.ApiError
	item.Desc = msg
	item.Err = data
	rawBytes := ioproto.GetBytesFromSerializer(serializer, item)
	return rawBytes
}

func ErrorBytes(msg string, serializer io.Serializer) []byte {
	return toErrorPool(str.UnsafeBytes(msg), serializer)
}

func ErrorStr(c *echo.Context, msg []byte) error {
	logger.Debug("converting error string to payload")
	logger.Error(str.UnsafeString(msg))
	rawBytes := toErrorPool(msg, c.ResponseSerializer())
	return c.FastBlob(protocol.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func ErrorWithMessage(c *echo.Context, code int, msg []byte, err error) error {
	logger.Debug("converting error with message to payload")
	logger.Error(err)
	rawBytes := toError(code, msg, str.UnsafeBytes(err.Error()), c.ResponseSerializer())
	return c.FastBlob(code, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func Error(c *echo.Context, err error) error {
	logger.Debug("converting error to payload")
	return ErrorStr(c, str.UnsafeBytes(err.Error()))
}

func ErrorCode(c *echo.Context, code int, err error) error {
	logger.Debug("converting error with code to error payload")
	logger.Error(err)
	rawBytes := toError(code, str.UnsafeBytes(err.Error()), nil, c.ResponseSerializer())
	return c.FastBlob(code, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}

func StackError(c *echo.Context, stackErr stack.Error) error {
	logger.Debug("converting stack error to error payload")
	logger.Error(stackErr)
	rawBytes := toError(protocol.StatusBadRequest, stackErr.Bytes(), nil, c.ResponseSerializer())
	return c.FastBlob(protocol.StatusBadRequest, echo.MIMEApplicationJSONCharsetUTF8, rawBytes)
}
