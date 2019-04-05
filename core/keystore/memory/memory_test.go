// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package memory

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/core/modules/cache"
)

func TestInMemoryKeyStorage_Set(t *testing.T) {
	type fields struct {
		cache *cache.MemoryCache
	}
	type args struct {
		key   string
		value WalletContent
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
			storage := &InMemoryKeyStorage{
				cache: tt.fields.cache,
			}
			storage.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestInMemoryKeyStorage_Get(t *testing.T) {
	type fields struct {
		cache *cache.MemoryCache
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   WalletContent
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := InMemoryKeyStorage{
				cache: tt.fields.cache,
			}
			got, got1 := storage.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryKeyStorage.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("InMemoryKeyStorage.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewInMemoryKeyStorage(t *testing.T) {
	tests := []struct {
		name string
		want *InMemoryKeyStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemoryKeyStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemoryKeyStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
