// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"github.com/zerjioang/etherniti/core/modules/cache"
	"github.com/zerjioang/etherniti/core/server"

	"github.com/zerjioang/etherniti/core/keystore/memory"
)

type WalletManager struct {
	wallet *memory.InMemoryKeyStorage
	cache  *cache.MemoryCache
}

func NewWalletManager() WalletManager {
	man := WalletManager{}
	man.wallet = memory.NewInMemoryKeyStorage()
	return man
}

// get token instance for given client and address
func InstantiateToken(cc *server.EthernitiContext, address string) (interface{}, error) {
	// get the client from cc context
	/*
		client, clientErr := ethclient.DialContext(context.Background(), "")
		if clientErr != nil {
			return nil, clientErr
		}
		instance, err := erc20.NewToken(address, client)
	*/
	//client := ethrpc.NewDefaultRPC("")
	//instance, err := erc20.NewToken(address, nil)
	// return instance, err
	return nil, nil
}

// proxy methods
func (wm WalletManager) Get(key []byte) (memory.WalletContent, bool) {
	return wm.wallet.Get(key)
}

func (wm *WalletManager) Set(key []byte, value memory.WalletContent) {
	wm.wallet.Set(key, value)
}
