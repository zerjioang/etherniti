package common

import (
	"github.com/zerjioang/etherniti/core/modules/stack"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type DatabaseObjectInterface interface {
	Key() []byte
	Value() []byte
	// creates new instance
	// to allow concurrent access, etc
	New() DatabaseObjectInterface
	CanRead(context *echo.Context, key string) error
	CanUpdate(context *echo.Context, key string) error
	CanDelete(context *echo.Context, key string) error
	CanWrite(context *echo.Context) error
	CanList(context *echo.Context) error
	// method used to decode http input byte data to go struct
	Bind(context *echo.Context) (DatabaseObjectInterface, stack.Error)
}
