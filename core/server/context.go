// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package server

import "github.com/labstack/echo"

/*
creating a custom context, allow us to add new features in a clean way
*/
type gaethwayContext struct {
	echo.Context
}

func NewgaethwayContext() gaethwayContext {
	ctx := gaethwayContext{}
	return ctx
}
