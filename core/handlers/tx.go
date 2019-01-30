// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	ctx = context.Background()
)

// eth transactions controller
type TransactionController struct {
}

// constructor like function
func NewTransactionController() TransactionController {
	ctl := TransactionController{}
	return ctl
}

// sends new eth transaction using given configuration
func (ctl EthController) Send() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Error(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Error(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Error(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

}

// implemented method from interface RouterRegistrable
func (ctl TransactionController) RegisterRouters(router *echo.Echo) {

}
