// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cache

import "github.com/zerjioang/etherniti/core/logger"

var (
	defaultCache *MemoryCache
)

func init() {
	logger.Debug("creating shared memory cache for etherniti proxy modules")
	defaultCache = NewMemoryCache()
}

type MemoryCache struct {
	c map[string]interface{}
}

func (cache MemoryCache) Get(key string) (interface{}, bool) {
	logger.Debug("reading value from shared memory cache")
	v, ok := cache.c[key]
	return v, ok
}

func (cache *MemoryCache) Set(key string, value interface{}) {
	logger.Debug("settings new value on shared memory cache")
	cache.c[key] = value
}

func NewMemoryCache() *MemoryCache {
	logger.Debug("creating new memory cache")
	m := new(MemoryCache)
	m.c = make(map[string]interface{}, 0)
	return m
}
