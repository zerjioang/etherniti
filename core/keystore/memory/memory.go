// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package memory

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/etherniti/core/logger"
)

var (
	emptyWallet = WalletContent{}
)

// in memory storage of accounts
type InMemoryKeyStorage struct {
	cache *cache.Cache
}

func (storage *InMemoryKeyStorage) Set(key string, value WalletContent) {
	logger.Info("adding new account to memory based wallet")
	storage.cache.Set(key, value, cache.DefaultExpiration)
}

func (storage InMemoryKeyStorage) Get(key string) (WalletContent, bool) {
	logger.Info("reading existing account from memory based wallet")
	raw, found := storage.cache.Get(key)
	if found {
		//cast
		content, ok := raw.(WalletContent)
		if ok {
			return content, true
		}
	}
	return emptyWallet, false
}

func NewInMemoryKeyStorage() *InMemoryKeyStorage {
	logger.Info("creating in-memory temporal wallet")
	s := new(InMemoryKeyStorage)
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	s.cache = cache.New(5*time.Minute, 10*time.Minute)
	return s
}
