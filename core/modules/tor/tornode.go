// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tor

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/hashset"
)

var (
	TornodeSet *hashset.HashSet
)

func init() {
	logger.Info("[module] loading tor nodes hash data")
	TornodeSet = hashset.NewHashSet()
	TornodeSet.LoadFromArray(tornodeList)
	TornodeSet.LoadFromArray(tornodeExitList)
}