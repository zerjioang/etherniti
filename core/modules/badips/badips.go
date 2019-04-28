// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package badips

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/hashset"
)

var (
	// https://www.howtoforge.com/tutorial/protect-your-server-computer-with-badips-and-fail2ban/
	// loaded from https://www.badips.com/get/list/any/5
	badIpList *hashset.HashSet
)

func init() {
	logger.Info("[module] loading bad ip list policy data")
	badIpList = hashset.NewHashSet()
	badIpList.LoadFromRaw(config.BadIpsFile, "\n")
}

func GetBadIPList() *hashset.HashSet {
	return badIpList
}

func IsBackListedIp(ip string) bool {
	return badIpList.Contains(ip)
}
