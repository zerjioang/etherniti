// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tor

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/hashset"
)

var (
	TornodeSet hashset.HashSetMutex
)

func init() {
	logger.Info("[module] loading tor nodes hash data")
	TornodeSet = hashset.NewHashSet()
	TornodeSet.LoadFromJsonArray(config.TorExitFile)
}
