// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package server

import "github.com/labstack/echo"

/*
creating a custom context, allow us to add new features in a clean way
*/
type MethWContext struct {
	echo.Context
}

func NewMethWContext() MethWContext {
	ctx := MethWContext{}
	return ctx
}
