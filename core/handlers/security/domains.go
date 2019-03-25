// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"encoding/json"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"io/ioutil"
)

// https://raw.githubusercontent.com/409H/EtherAddressLookup/master/blacklists/domains.json
// last update: mar-23-2019
type DomainBlacklistModel []string
var (
	domainBlacklist DomainBlacklistModel
	//domain blacklist as bytes
	domainBlacklistBytes []byte
)

// load blacklist information
func init(){
	logger.Debug("loading blacklist information")
	data, err := ioutil.ReadFile(config.BlacklistedDomainFile)
	if err != nil {
		logger.Error("could not read blacklist data")
		return
	}
	domainBlacklistBytes = data
	json.Unmarshal(data, &domainBlacklist)
}

func DomainBlacklistBytesData() []byte{
	return domainBlacklistBytes
}