// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package bots

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/hashset"
)

var (
	badBotsList hashset.HashSetMutex
)

func init() {
	logger.Info("[module] loading anti-bots policy data")
	badBotsList = hashset.NewHashSet()
	badBotsList.LoadFromJsonArray(config.AntiBotsFile)
}
