// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package server

import "github.com/labstack/echo"

// creating a custom context,
// allow us to add new features in a clean way
type GaethwayContext struct {
	echo.Context
}

func NewgaethwayContext() GaethwayContext {
	ctx := GaethwayContext{}
	return ctx
}
