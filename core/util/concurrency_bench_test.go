// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	mut *sync.RWMutex
	endpoint string
}

func BenchmarkPMutexSet(b *testing.B) {
	config := Config{}
	config.mut = new(sync.RWMutex)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.mut.Lock()
			config.endpoint = "api.example.com"
			config.mut.Unlock()
		}
	})
}

func BenchmarkPMutexGet(b *testing.B) {
	config := Config{endpoint: "api.example.com"}
	config.mut = new(sync.RWMutex)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.mut.RLock()
			_ = config.endpoint
			config.mut.RUnlock()
		}
	})
}

func BenchmarkPAtomicSet(b *testing.B) {
	var config = new(atomic.Value)
	c := Config{endpoint: "api.example.com"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Store(c)
		}
	})
}

func BenchmarkPAtomicGet(b *testing.B) {
	var config = new(atomic.Value)
	config.Store(Config{endpoint: "api.example.com"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = config.Load().(Config)
		}
	})
}
