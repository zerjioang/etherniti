// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/labstack/echo"
	"github.com/zerjioang/gaethway/core/api"
	"github.com/zerjioang/gaethway/core/eth"
	"github.com/zerjioang/gaethway/core/eth/rpc"
	"github.com/zerjioang/gaethway/core/modules/ethfork/log"
	"github.com/zerjioang/gaethway/core/util"
	"net/http"
)

// ganache controller
type GanacheController struct {
	// in memory based wallet manager
	walletManager eth.WalletManager
}

// constructor like function
func NewGanacheController(manager eth.WalletManager) GanacheController {
	ctl := GanacheController{}
	ctl.walletManager = manager
	return ctl
}

func (ctl GanacheController) getAccounts(c echo.Context) error {
	client := ethrpc.NewDefaultRPC("http://127.0.0.1:8545")
	list, err := client.EthAccounts()
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusInternalServerError,
			util.GetJsonBytes(
				api.NewApiError(http.StatusInternalServerError, err.Error()),
			),
		)
	} else {
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				api.NewApiResponse("ethereum accounts readed", list),
			),
		)
	}
}
/*
{
  "jsonrpc": "2.0",
  "method": "eth_getBalance",
  "params": ["0x0ADfCCa4B2a1132F82488546AcA086D7E24EA324", "latest"],
  "id": 1
}
 */
func (ctl GanacheController) getAccountsWithBalance(c echo.Context) error {
	client := ethrpc.NewDefaultRPC("http://127.0.0.1:8545")
	list, err := client.EthAccounts()

	type wrapper struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
		Eth string `json:"eth"`
		Key string `json:"key"`
	}
	wrapperList := make([]wrapper, len(list))

	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusInternalServerError,
			util.GetJsonBytes(
				api.NewApiError(http.StatusInternalServerError, err.Error()),
			),
		)
	} else {
		//iterate over account
		for i:=0; i<len(list);i++{
			currentAccount := list[i]
			bigInt, err := client.EthGetBalance(currentAccount, "latest")
			if err != nil {
				log.Error("failed to get account balance", currentAccount, err)
			} else {
				item := &wrapperList[i]
				item.Account = currentAccount
				item.Balance = bigInt.String()
				item.Eth = eth.ToEth(bigInt).String()
				item.Key = ""
			}
		}
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				api.NewApiResponse("ethereum accounts and their balance readed", wrapperList),
			),
		)
	}
}

func (ctl GanacheController) getBlocks(c echo.Context) error {
	return nil
}

func (ctl GanacheController) getTransactions(c echo.Context) error {
	return nil
}

// implemented method from interface RouterRegistrable
func (ctl GanacheController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing ganache controller methods")
	//http://localhost:8080/eth/create
	router.GET("/v1/ganache/accounts", ctl.getAccounts)
	router.GET("/v1/ganache/accountsBalanced", ctl.getAccountsWithBalance)
	router.GET("/v1/ganache/blocks", ctl.getBlocks)
	router.GET("/v1/ganache/tx", ctl.getTransactions)
}
