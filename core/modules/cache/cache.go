// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cache

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/concurrentmap"
)

var (
	defaultCache *MemoryCache
)

func init() {
	defaultCache = NewMemoryCache()
}

type MemoryCache struct {
	c concurrentmap.ConcurrentMap
}

func (cache MemoryCache) Get(key string) (interface{}, bool) {
	logger.Debug("reading value from global memory cache")
	v, ok := cache.c.Get(key)
	return v, ok
}

func (cache *MemoryCache) Set(key string, value interface{}) {
	logger.Debug("settings new value on global memory cache")
	cache.c.Set(key, value)
}

func NewMemoryCache() *MemoryCache {
	logger.Debug("creating new global cache")
	m := new(MemoryCache)
	m.c = concurrentmap.New()
	return m
}

func Instance() *MemoryCache {
	return defaultCache
}
