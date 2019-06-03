// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tokenlist

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tidwall/gjson"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
)

// token data is loaded from json file downloaded from
// https://raw.githubusercontent.com/kvhnuke/etherwallet/mercury/app/scripts/tokens/ethTokens.json
// last update: mar-23-2019
type TokenInfo struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
	Type    string `json:"type"`
	Decimal int    `json:"decimal"`
}
// list containing TokenInfo
// this element is stored in heap
type TokenInfoList []TokenInfo

var (
	tokenlistData  TokenInfoList
	tokenlistBytes []byte
)

// load token data information
func init() {
	logger.Debug("loading token list information")
	data, err := ioutil.ReadFile(config.TokenListFile)
	if err != nil {
		logger.Error("could not read token list data")
		return
	}
	tokenlistBytes = data
	unErr := json.Unmarshal(tokenlistBytes, &tokenlistData)
	if unErr != nil {
		logger.Error("could not unmarshal token list data")
		return
	}
}

// fetch token data (address only) by token name
// todo implement some caching mecanism
func GetTokenAddressByName(name string) string {
	value := gjson.GetBytes(tokenlistBytes, `#[symbol=="`+name+`"].address`)
	return value.Str
}
// fetch token data (symbol only) by token address
// todo implement some caching mecanism
func GetTokenSymbol(address string) string {
	value := gjson.GetBytes(tokenlistBytes, `#[address=="`+address+`"].symbol`)
	return value.Str
}
