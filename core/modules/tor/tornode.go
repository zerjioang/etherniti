// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tor

import (
	"io/ioutil"
	"strings"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/hashset"
	"github.com/zerjioang/etherniti/core/util/ip"
	"github.com/zerjioang/etherniti/core/util/str"
)

type TorList struct {
	hashset.HashSetAtomic
}

func (l *TorList) LoadIps(path string) {
	logger.Debug("loading tor node list")
	if path != "" {
		logger.Debug("loading tor list with raw data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		itemList := strings.Split(str.UnsafeString(data), "\n")
		if itemList != nil {
			for _, v := range itemList {
				//convert string ip to uint
				ipvalue := ip.Ip2intLow(v)
				l.UnsafeAddUint32(ipvalue)
			}
		}
	}
}

var (
	TornodeSet TorList
)

func init() {
	logger.Info("[module] loading tor nodes hash data")
	TornodeSet.LoadIps(config.TorExitFile)
}
