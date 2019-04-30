// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"github.com/zerjioang/etherniti/core/keystore/memory"
	"github.com/zerjioang/etherniti/core/modules/cache"
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

// proxy methods
func (wm WalletManager) Get(key []byte) (memory.WalletContent, bool) {
	return wm.wallet.Get(key)
}

func (wm *WalletManager) Set(key []byte, value memory.WalletContent) {
	wm.wallet.Set(key, value)
}
