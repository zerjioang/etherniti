// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/api"

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
	fuzzyData []byte
	responseName = []byte("domain analyzed")
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
	blackData = str.GetJsonBytes(pm.Blacklist)
	whiteData = str.GetJsonBytes(pm.Whitelist)
	fuzzyData = str.GetJsonBytes(pm.Fuzzylist)
}

func PhishingBlacklistRawBytes() []byte {
	return blackData
}
func PhishingWhitelistRawBytes() []byte {
	return whiteData
}

func FuzzyDataRawBytes() []byte {
	return fuzzyData
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func IsDangerousDomain(domain string) []byte {
	type response struct {
		Domain string `json:"domain"`
		Trust  bool   `json:"trust"`
		//metadata
		Title   string `json:"title"`
		Message string `json:"message"`
	}
	warn := contains(pm.Blacklist, domain) || contains(DomainBlacklist(), domain)
	if warn {
		responseData := response{
			Title:   "Deceptive domain detected",
			Domain:  domain,
			Message: "The domain you requested has been identified as being potentially problematic. This could be because a user has reported a problem, a black-list service reported a problem, or because we have detected potentially malicious content.",
			Trust:   false,
		}
		return api.ToSuccess(responseName, responseData)
	} else {
		responseData := response{
			Title:   "Clean domain detected",
			Domain:  domain,
			Message: "The domain you requested has not been blacklisted.",
			Trust:   true,
		}
		return api.ToSuccess(responseName, responseData)
	}
}
