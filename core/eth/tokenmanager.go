// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/gaethway/core/keystore/memory"
	"github.com/zerjioang/gaethway/core/modules/token"
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
func InstantiateToken(client *ethclient.Client, address common.Address) (*token.Token, error) {
	instance, err := token.NewToken(address, client)
	return instance, err
}

// proxy methods
func (wm WalletManager) Get(key string) (memory.WalletContent, bool) {
	return wm.wallet.Get(key)
}

func (wm *WalletManager) Set(key string, value memory.WalletContent) {
	wm.wallet.Set(key, value)
}
