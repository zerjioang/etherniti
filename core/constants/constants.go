// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package constants

const (
	// set system pointer size
	PointerSize = 32 + int(^uintptr(0)>>63<<5)
	// api version
	ApiVersion = "/v1"
	// context free api path
	RootApi = "/"
	// context dependant api path
	PrivateApi = "/private"
)
