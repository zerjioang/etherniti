// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build appengine

package log

import (
	"io"
	"os"
)

func output() io.Writer {
	return os.Stdout
}
