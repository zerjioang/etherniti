// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zerjioang/etherniti/core/server"
	"time"

	"github.com/zerjioang/etherniti/core/modules/token/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/etherniti/core/keystore/memory"
)

type WalletManager struct {
	wallet *memory.InMemoryKeyStorage
	cache  *cache.Cache
}

func NewWalletManager() WalletManager {
	man := WalletManager{}
	man.wallet = memory.NewInMemoryKeyStorage()
	man.cache = cache.New(5*time.Minute, 10*time.Minute)
	return man
}

// get token instance for given client and address
func InstantiateToken(cc *server.EthernitiContext, address common.Address) (*erc20.ERC20Token, error) {
	// get the client from cc context
	client, clientErr := ethclient.DialContext(context.Background(), "")
	if clientErr != nil {
		return nil, clientErr
	}
	instance, err := erc20.NewToken(address, client)
	return instance, err
}

// proxy methods
func (wm WalletManager) Get(key string) (memory.WalletContent, bool) {
	return wm.wallet.Get(key)
}

func (wm *WalletManager) Set(key string, value memory.WalletContent) {
	wm.wallet.Set(key, value)
}
