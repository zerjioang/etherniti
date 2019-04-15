// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build pre prod

package api

import "github.com/zerjioang/etherniti/core/modules/hashset"

func GetBadBotsList() *hashset.HashSet {
	return badBotsList
}
