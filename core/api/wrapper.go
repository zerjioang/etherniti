// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"sync"

	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"

	"github.com/zerjioang/go-hpc/lib/codes"

	"github.com/zerjioang/go-hpc/common"

	"github.com/zerjioang/go-hpc/util/str"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/lib/stack"
)

const (
	none = ""
)

var (
	errorPool   sync.Pool
	successPool sync.Pool
)

func init() {
	logger.Debug("loading api wrapper data")
	errorPool = sync.Pool{
		New: func() interface{} {
			return dto.NewApiError(codes.StatusBadRequest, none)
		},
	}
	successPool = sync.Pool{
		New: func() interface{} {
			return dto.NewApiResponse([]byte{}, nil)
		},
	}
}

// return success response to client context
func SendSuccess(c *shared.EthernitiContext, logMsg []byte, response interface{}) error {
	return c.FastBlob(
		codes.StatusOK,
		c.ResponseContentType(),
		ToSuccess(logMsg, response, c.ResponseSerializer()),
	)
}

// return success response to client context
func SendRawSuccess(c *shared.EthernitiContext, content []byte) error {
	return c.FastBlob(
		codes.StatusOK,
		c.ResponseContentType(),
		content,
	)
}

func SendSuccessPool(c *shared.EthernitiContext, logMsg []byte, v interface{}) error {
	item := successPool.Get().(*dto.ApiResponse)
	item.Message = str.UnsafeString(logMsg)
	item.Data = v
	rawBytes := encoding.GetBytesFromSerializer(c.ResponseSerializer(), item)
	// put item back to the pool
	successPool.Put(item)

	return c.FastBlob(
		codes.StatusOK,
		c.ResponseContentType(),
		rawBytes,
	)
}

// return success blob response to client context
func SendSuccessBlob(c *shared.EthernitiContext, raw []byte) error {
	return c.FastBlob(codes.StatusOK, c.ResponseContentType(), raw)
}

func Success(c *shared.EthernitiContext, msg []byte, result []byte) error {
	item := successPool.Get().(*dto.ApiResponse)
	item.Message = str.UnsafeString(msg)
	item.Data = result
	rawBytes := encoding.GetBytesFromSerializer(c.ResponseSerializer(), item)
	// put item back to the pool
	successPool.Put(item)
	return c.FastBlob(codes.StatusOK, c.ResponseContentType(), rawBytes)
}

func ToSuccess(msg []byte, result interface{}, serializer common.Serializer) []byte {
	var item dto.ApiResponse
	item.Message = str.UnsafeString(msg)
	item.Data = result
	rawBytes := encoding.GetBytesFromSerializer(serializer, item)
	// put item back to the pool
	return rawBytes
}

func toErrorPool(msg string, serializer common.Serializer) []byte {
	//get item from pool
	item := errorPool.Get().(*dto.ApiError)
	item.Desc = msg
	rawBytes := encoding.GetBytesFromSerializer(serializer, item)
	// put item back to the pool
	errorPool.Put(item)
	return rawBytes
}

func toError(code codes.HttpStatusCode, msg string, data string, serializer common.Serializer) []byte {
	var item dto.ApiError
	item.Desc = msg
	item.Err = data
	rawBytes := encoding.GetBytesFromSerializer(serializer, item)
	return rawBytes
}

func ErrorStr(c *shared.EthernitiContext, msg string) error {
	return ErrorBytes(c, str.UnsafeBytes(msg))
}

func Error(c *shared.EthernitiContext, err error) error {
	rawBytes := toError(codes.StatusBadRequest, "", err.Error(), c.ResponseSerializer())
	return ErrorBytes(c, rawBytes)
}

func ErrorBytes(c *shared.EthernitiContext, msg []byte) error {
	return ErrorBytesWithCode(c, codes.StatusBadRequest, msg)
}

func ErrorBytesWithCode(c *shared.EthernitiContext, code codes.HttpStatusCode, msg []byte) error {
	return returnError(c, code, msg)
}

func ErrorWithMessage(c *shared.EthernitiContext, code codes.HttpStatusCode, msg []byte, err error) error {
	msgstr := str.UnsafeString(msg)
	rawBytes := toError(code, msgstr, err.Error(), c.ResponseSerializer())
	return returnError(c, code, rawBytes)
}

func ErrorCode(c *shared.EthernitiContext, code codes.HttpStatusCode, err error) error {
	rawBytes := toError(code, err.Error(), none, c.ResponseSerializer())
	return returnError(c, code, rawBytes)
}

func StackError(c *shared.EthernitiContext, stackErr stack.Error) error {
	rawBytes := toError(codes.StatusBadRequest, stackErr.Error(), none, c.ResponseSerializer())
	return returnError(c, codes.StatusBadRequest, rawBytes)
}

func returnError(c *shared.EthernitiContext, code codes.HttpStatusCode, rawBytes []byte) error {
	c.Response().Header().Set("Connection", "Close")
	// use http client requested encoding as response content type
	ct := c.ResponseContentType()
	c.Response().Header().Set("Content-Type", ct.String())
	encodedBytes := encoding.GetBytesFromMode(ct, rawBytes)
	return c.FastBlob(code, ct, encodedBytes)
}

func Redirect(c *shared.EthernitiContext, url string) error {
	return c.Redirect(codes.StatusTemporaryRedirect, url)
}
