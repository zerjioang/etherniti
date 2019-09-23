// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !appengine

package log

import (
	"io"

	"github.com/mattn/go-colorable"
)

var (
	colorableOut = colorable.NewColorableStdout()
)
func output() io.Writer {
	return colorableOut
}
