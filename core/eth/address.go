// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"context"
	"regexp"

	"github.com/zerjioang/etherniti/core/eth/fixtures"

	"github.com/zerjioang/etherniti/core/eth/rpc"
)

var (
	addressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	ctx          = context.Background()
)

const (
	LatestBlockNumber = "latest"
)

// converts an string ethereum address to eth address
// addr: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
func ConvertAddress(addr string) fixtures.Address {
	return fixtures.HexToAddress(addr)
}

// check if an address is syntactically valid or not
func IsValidAddress(addr string) bool {
	return addressRegex.MatchString(addr)
}

func IsSmartContractAddress(client ethrpc.EthRPC, addr string) (bool, error) {
	bytecode, err := client.EthGetCode(addr, LatestBlockNumber) // empty is latest
	// if the address has valid bytecode, is a contract
	// if is not code addres 0x is returned
	return len(bytecode) > 2, err
}
