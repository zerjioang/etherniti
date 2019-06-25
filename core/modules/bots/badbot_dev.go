// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build dev !dev
// +build !pre
// +build !prod

package bots

import "github.com/zerjioang/etherniti/core/modules/hashset"

func init() {
	badBotsList.Remove("apachebench")
}

func GetBadBotsList() hashset.HashSetWORM {
	return badBotsList
}
