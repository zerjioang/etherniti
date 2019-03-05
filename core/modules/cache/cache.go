// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cache

import (
	"time"

	"github.com/allegro/bigcache"
)

var (
	defaultCacheConfig = bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute, //eviction
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// prints information about additional memory allocation
		Verbose: true,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,
		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,
		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}
)

type MemoryCache struct {
	cache *bigcache.BigCache
}

func (cache MemoryCache) Get(key string) (interface{}, bool) {
	return nil, false
}

func (cache MemoryCache) Set(key string, value interface{}, duration time.Duration) {

}

func NewMemoryCache() *MemoryCache {
	m := new(MemoryCache)
	m.cache, _ = bigcache.NewBigCache(defaultCacheConfig)
	return m
}
