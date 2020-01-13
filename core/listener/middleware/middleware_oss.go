// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// edition tags: opensource
// +build oss

package middleware

import (
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

// this is open-source edition middleware
func secure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ApplyDefaultCommonHeaders(c)
		ApplyDefaultSecurityHeaders(c)
		return next(c)
	}
}
