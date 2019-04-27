// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !dev
// +build !pre
// +build prod

package bots

import "github.com/zerjioang/etherniti/core/modules/hashset"

func GetBadBotsList() *hashset.HashSet {
	return badBotsList
}