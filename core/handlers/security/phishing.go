// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"encoding/json"
	"io/ioutil"
	"unsafe"

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
	json.Unmarshal(data, &pm)
	blackData = *(*[]byte)(unsafe.Pointer(&pm.Blacklist))
	whiteData = *(*[]byte)(unsafe.Pointer(&pm.Whitelist))
}

func PhishingBlacklistRawBytes() []byte {
	return blackData
}
func PhishingWhitelistRawBytes() []byte {
	return whiteData
}
