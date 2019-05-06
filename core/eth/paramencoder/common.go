// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package paramencoder

import (
	"github.com/zerjioang/etherniti/core/eth/fixtures/abi"
	"github.com/zerjioang/etherniti/core/eth/fixtures/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/encoding/hex"
)

func GeneratePayload(aibmodel *abi.ABI, functionName string) (string, error) {
	//preload symbol function params
	temp, err := aibmodel.Pack(functionName)
	if err != nil {
		logger.Error("failed to load ERC20 '", functionName, "' interaction model")
		return "", err
	} else {
		// add 32 byte padding to set that this function has no parameters
		data := hex.ToEthHex(common.RightPadBytes(temp, 32+len(temp)))
		return data, err
	}
}
