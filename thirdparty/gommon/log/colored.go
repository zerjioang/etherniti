// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !appengine

package log

import (
	"io"

	"github.com/mattn/go-colorable"
)

func output() io.Writer {
	return colorable.NewColorableStdout()
}
