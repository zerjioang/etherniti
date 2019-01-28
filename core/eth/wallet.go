// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
)

// content stored in the wallet
type WalletContent struct {
	// address of the account
	ethAddress common.Address
	// account's private key enconded as ecdsa go struct
	privateKey ecdsa.PrivateKey
	// client interaction eth/quorum node for processing interactions
	connectionClient *ethclient.Client
}

// returns the client linked to the saved wallet
func (wallet WalletContent) Client() *ethclient.Client {
	return wallet.connectionClient
}
