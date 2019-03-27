// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"encoding/json"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/util"
	"io/ioutil"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
)

// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json
// last update: mar-23-2019

//phishing data loader model
type PhishingModel struct {
	Version   int      `json:"version"`
	Tolerance int      `json:"tolerance"`
	Fuzzylist []string `json:"fuzzylist"`
	Whitelist []string `json:"whitelist"`
	Blacklist []string `json:"blacklist"`
}

var (
	pm        PhishingModel
	whiteData []byte
	blackData []byte
)

func init() {
	logger.Debug("loading phising model information")
	data, err := ioutil.ReadFile(config.PhishingDomainFile)
	if err != nil {
		logger.Error("could not read phising model data")
		return
	}
	unmarshalErr := json.Unmarshal(data, &pm)
	if unmarshalErr != nil {
		logger.Error("could not unmarshal phising model data")
		return
	}
	blackData = util.GetJsonBytes(pm.Blacklist)
	whiteData = util.GetJsonBytes(pm.Whitelist)
}

func PhishingBlacklistRawBytes() []byte {
	return blackData
}
func PhishingWhitelistRawBytes() []byte {
	return whiteData
}

func IsDangerousDomain(domain string) []byte {
	type response struct {
		Domain string `json:"domain"`
		Trust  bool   `json:"trust"`
		Score  uint8  `json:"score"`
		//metadata
		Title   string `json:"title"`
		Message string `json:"message"`
	}
	responseData := response{
		Title:   "Deceptive domain detected",
		Domain:  domain,
		Message: "The domain you requested has been identified as being potentially problematic. This could be because a user has reported a problem, a black-list service reported a problem, or because we have detected potentially malicious content.",
		Score:   uint8(7),
		Trust:   false,
	}
	return api.ToSuccess("domain verified", responseData)
}
