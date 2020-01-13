package wrap

import (
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

/*
this functions wraps the default context to
etherniti context received as parameter to
all controllers
*/

type extendedHandlerFunc func(c *shared.EthernitiContext) error

func Call(handler extendedHandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ec, ok := c.(*shared.EthernitiContext)
		if ok && ec != nil {
			return handler(ec)
		}
		return nil
	}
}
