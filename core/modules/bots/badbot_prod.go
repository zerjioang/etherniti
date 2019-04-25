// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build prod
// +build !pre
// +build !dev

package bots

import "github.com/zerjioang/etherniti/core/modules/hashset"

func GetBadBotsList() *hashset.HashSet {
	return badBotsList
}
