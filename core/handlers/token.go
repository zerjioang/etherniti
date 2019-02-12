// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"fmt"
	"math"
	"math/big"

	"github.com/zerjioang/etherniti/core/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/eth"
)

// token controller
type TokenController struct {
	// in memory based wallet manager
	walletManager eth.WalletManager
}

// constructor like function
func NewTokenController(manager eth.WalletManager) TokenController {
	ctl := TokenController{}
	ctl.walletManager = manager
	return ctl
}

func (ctl TokenController) instantiate(c echo.Context) error {
	targetAddr := c.Param("address")
	ethAddr := eth.ConvertAddress(targetAddr)
	client, err := ctl.getClientInstance(c)
	if err == nil {
		instance, err := eth.InstantiateToken(client, ethAddr)
		if err == nil && instance != nil {
			//todo save token instance in memory
		}
	}
	return err
}

func (ctl TokenController) summary(c echo.Context) error {
	targetAddr := c.Param("address")
	ethAddr := eth.ConvertAddress(targetAddr)
	client, err := ctl.getClientInstance(c)
	if err == nil {
		instance, err := eth.InstantiateToken(client, ethAddr)
		if err == nil && instance != nil {
			//todo save token instance in memory

			//show token summary
			bal, err := instance.BalanceOf(&bind.CallOpts{}, ethAddr)
			if err != nil {
				log.Fatal(err)
			}

			name, err := instance.Name(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			symbol, err := instance.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			decimals, err := instance.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("name: %s\n", name)         // "name: Golem Network"
			fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
			fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

			fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

			fbal := new(big.Float)
			fbal.SetString(bal.String())
			value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

			fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
		}
	}
	return err
}

// from incoming http request, it recovers the eth client linked to it
func (ctl TokenController) getClientInstance(c echo.Context) (*ethclient.Client, error) {
	requestProfileKey := c.Request().Header.Get(config.HttpProfileHeaderkey)
	wallet, found := ctl.walletManager.Get(requestProfileKey)
	if !found {
		return nil, errNoConnectionProfile
	}
	return wallet.Client(), nil
}

// implemented method from interface RouterRegistrable
func (ctl TokenController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing eth controller methods")
	//http://localhost:8080/eth/create
	router.POST("/v1/token/instance", ctl.instantiate)
}
