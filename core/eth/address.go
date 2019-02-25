// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"context"
	"math"
	"math/big"
	"regexp"

	"github.com/zerjioang/etherniti/core/eth/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	addressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	ctx          = context.Background()
)

// converts an string ethereum address to eth address
// addr: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
func ConvertAddress(addr string) common.Address {
	return common.HexToAddress(addr)
}

// check if an address is syntactically valid or not
func IsValidAddress(addr string) bool {
	return addressRegex.MatchString(addr)
}

func IsSmartContractAddress(client ethrpc.EthRPC, addr string) (bool, error) {
	bytecode, err := client.EthGetCode(addr, "") // empty is latest
	// if the address has valid bytecode, is a contract
	return len(bytecode) > 0, err
}

/*
fmt.Println(address.Hex())        // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
fmt.Println(address.Bytes())      // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
*/

func GetAccountBalance(client *ethclient.Client, addr common.Address) (*big.Int, error) {
	return client.BalanceAt(ctx, addr, nil)
}

func GetAccountBalanceAtBlock(client *ethclient.Client, addr common.Address, blockNumber *big.Int) (*big.Int, error) {
	return client.BalanceAt(ctx, addr, blockNumber)
}

func ToEth(balance big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}
