// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package constants

import "os"

// database config
var (
	Home             = os.Getenv("HOME")
	DatabaseRootPath = Home + "/.etherniti/"
)
