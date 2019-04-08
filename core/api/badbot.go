// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/hashset"
)

var (
	BadBotsList *hashset.HashSet
)

func init() {
	logger.Info("[module] loading anti-bots policy data")
	BadBotsList = hashset.NewHashSet()
	BadBotsList.LoadFromArray(badBotsList)
}
