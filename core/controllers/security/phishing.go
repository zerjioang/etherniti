// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto"
	"github.com/zerjioang/etherniti/shared/protocol/io"

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
	pm           PhishingModel
	whiteData    []byte
	blackData    []byte
	fuzzyData    []byte
	responseName = []byte("domain analyzed")
	//todo remove this in next releases. manage serializer on demand
	defaultSerializer, _ = ioproto.EncodingModeSelector(io.ModeJson)
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
	blackData = ioproto.GetBytesFromSerializer(defaultSerializer, pm.Blacklist)
	whiteData = ioproto.GetBytesFromSerializer(defaultSerializer, pm.Whitelist)
	fuzzyData = ioproto.GetBytesFromSerializer(defaultSerializer, pm.Fuzzylist)
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

type response struct {
	Domain string `json:"domain"`
	Trust  bool   `json:"trust"`
	//metadata
	Title   string `json:"title"`
	Message string `json:"message"`
}

func isDangerous(domain string) bool {
	return contains(pm.Blacklist, domain) || contains(DomainBlacklist(), domain)
}
func IsDangerousDomain(domain string) []byte {
	warn := isDangerous(domain)
	if warn {
		responseData := response{
			Title:   "deceptive domain detected",
			Domain:  domain,
			Message: "the domain you requested has been identified as being potentially problematic. This could be because a user has reported a problem, a black-list service reported a problem, or because we have detected potentially malicious content.",
			Trust:   false,
		}
		return api.ToSuccess(responseName, responseData, defaultSerializer)
	} else {
		responseData := response{
			Title:   "clean domain detected",
			Domain:  domain,
			Message: "the domain you requested has not been blacklisted.",
			Trust:   true,
		}
		return api.ToSuccess(responseName, responseData, defaultSerializer)
	}
}
