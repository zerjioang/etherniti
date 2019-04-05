// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cache

import (
	"reflect"
	"testing"
	"time"

	"github.com/allegro/bigcache"
)

func TestMemoryCache_Get(t *testing.T) {
	type fields struct {
		cache *bigcache.BigCache
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := MemoryCache{
				cache: tt.fields.cache,
			}
			got, got1 := cache.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoryCache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MemoryCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMemoryCache_Set(t *testing.T) {
	type fields struct {
		cache *bigcache.BigCache
	}
	type args struct {
		key      string
		value    interface{}
		duration time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := MemoryCache{
				cache: tt.fields.cache,
			}
			cache.Set(tt.args.key, tt.args.value, tt.args.duration)
		})
	}
}

func TestNewMemoryCache(t *testing.T) {
	tests := []struct {
		name string
		want *MemoryCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
