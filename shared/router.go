// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package shared

import "github.com/labstack/echo"

type AutoRouteable interface {
	RegisterRouters(router *echo.Echo)
}
