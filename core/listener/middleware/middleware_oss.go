// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// edition tags: opensource
// +build oss

package middleware

import (
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// this is open-source edition middleware
func secure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ApplyDefaultSecurityHeaders(c)
		return next(c)
	}
}
