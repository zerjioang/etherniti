// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package memory

import (
	"crypto/ecdsa"
	"github.com/zerjioang/etherniti/core/eth/rpc"

	"github.com/ethereum/go-ethereum/common"
)

// content stored in the wallet
type WalletContent struct {
	// address of the account
	ethAddress common.Address
	// account's private key enconded as ecdsa go struct
	privateKey ecdsa.PrivateKey
	// client interaction eth/quorum node for processing interactions
	connectionClient ethrpc.EthRPC
}

// returns the client linked to the saved wallet
func (wallet WalletContent) Client() ethrpc.EthRPC {
	return wallet.connectionClient
}
