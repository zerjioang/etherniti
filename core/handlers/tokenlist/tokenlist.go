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

func init() {
	// token data is loaded from json file downloaded from
	// https://raw.githubusercontent.com/kvhnuke/etherwallet/mercury/app/scripts/tokens/ethTokens.json
}

// https://raw.githubusercontent.com/409H/EtherAddressLookup/master/blacklists/domains.json
// last update: mar-23-2019
type TokenData struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
	Type    string `json:"type"`
	Decimal int    `json:"decimal"`
}
type TokenListModel []TokenData

var (
	tokenlistData  TokenListModel
	tokenlistBytes []byte
)

// load blacklist information
func init() {
	logger.Debug("loading loken list information")
	data, err := ioutil.ReadFile(config.TokenListFile)
	if err != nil {
		logger.Error("could not read tokenlist data")
		return
	}
	tokenlistBytes = data
	unErr := json.Unmarshal(tokenlistBytes, &tokenlistData)
	if unErr != nil {
		logger.Error("could not unmarshal token list data")
		return
	}
}

func GetTokenAddressByName(name string) string {
	value := gjson.GetBytes(tokenlistBytes, `#[symbol=="`+name+`"].address`)
	return value.Str
}

func GetTokenSymbol(address string) string {
	value := gjson.GetBytes(tokenlistBytes, `#[address=="`+address+`"].symbol`)
	return value.Str
}
