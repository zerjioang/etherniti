// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package middleware

import (
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type (
	// Skipper defines a function to skip middleware. Returning true skips processing
	// the middleware.
	Skipper func(echo.ContextInterface) bool

	// BeforeFunc defines a function which is executed just before the middleware.
	BeforeFunc func(echo.ContextInterface)
)

// DefaultSkipper returns false which processes the middleware.
func DefaultSkipper(echo.ContextInterface) bool {
	return false
}