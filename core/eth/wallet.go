// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

// content stored in the wallet
type WalletContent struct {
	ethAddress common.Address
	privateKey *ecdsa.PrivateKey
}
