// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cache

var (
	defaultCache *MemoryCache
)

func init(){
	defaultCache = NewMemoryCache()
}
type MemoryCache struct {
	c map[string]interface{}
}

func (cache MemoryCache) Get(key string) (interface{}, bool) {
	v, ok := cache.c[key]
	return v, ok
}

func (cache *MemoryCache) Set(key string, value interface{}) {
	cache.c[key] = value
}

func NewMemoryCache() *MemoryCache {
	m := new(MemoryCache)
	m.c = make(map[string]interface{}, 0)
	return m
}
