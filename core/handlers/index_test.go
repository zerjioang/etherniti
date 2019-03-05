// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"testing"
	"time"

	"github.com/zerjioang/etherniti/core/listener/base"

	"github.com/pkg/profile"
)

func TestIndexConcurrency(t *testing.T) {
	t.Run("index-single-echo", func(t *testing.T) {
		testServer := base.NewServer()
		ctx := NewContext(testServer)
		err := Index(ctx)
		if err != nil {
			t.Log(err)
		}
	})
	t.Run("index-concurrency-echo", func(t *testing.T) {
		times := 100
		testServer := NewServer()
		for i := 0; i < times; i++ {
			go func() {
				ctx := NewContext(testServer)
				err := Index(ctx)
				if err != nil {
					t.Log(err)
				}
			}()
		}
		time.Sleep(2 * time.Second)
	})

	t.Run("status-single-echo", func(t *testing.T) {
		testServer := NewServer()
		ctl := NewIndexController()
		ctx := NewContext(testServer)
		err := ctl.Status(ctx)
		if err != nil {
			t.Log(err)
		}
	})
	t.Run("status-concurrency-echo", func(t *testing.T) {
		times := 100
		testServer := NewServer()
		ctl := NewIndexController()
		for i := 0; i < times; i++ {
			go func() {
				ctx := NewContext(testServer)
				err := ctl.Status(ctx)
				if err != nil {
					t.Log(err)
				}
			}()
		}
		time.Sleep(2 * time.Second)
	})
	t.Run("status-concurrency-controller", func(t *testing.T) {
		times := 100
		ctl := NewIndexController()
		for i := 0; i < times; i++ {
			go func() {
				result := ctl.status()
				if result == nil || len(result) == 0 {
					t.Log("invalid status result")
				}
			}()
		}
		time.Sleep(2 * time.Second)
	})
	t.Run("status-single-controller", func(t *testing.T) {
		ctl := NewIndexController()
		result := ctl.status()
		if result == nil || len(result) == 0 {
			t.Log("invalid status result")
		}
	})
}

func TestIndexProfiling(t *testing.T) {
	t.Run("status", func(t *testing.T) {
		t.Run("cpu", func(t *testing.T) {
			// go tool pprof --pdf ~/go/bin/yourbinary /var/path/to/cpu.pprof > file.pdf
			defer profile.Start().Stop()
			ctl := NewIndexController()
			for n := 0; n < 1000000; n++ {
				ctl.status()
			}
		})
		t.Run("mem", func(t *testing.T) {
			// go tool pprof --pdf ~/go/bin/yourbinary /var/path/to/cpu.pprof > file.pdf
			defer profile.Start(profile.MemProfile).Stop()
			ctl := NewIndexController()
			for n := 0; n < 1000000; n++ {
				ctl.status()
			}
		})
	})
	t.Run("integrity", func(t *testing.T) {
		t.Run("cpu", func(t *testing.T) {
			// go tool pprof --pdf ~/go/bin/yourbinary /var/path/to/cpu.pprof > file.pdf
			defer profile.Start().Stop()
			ctl := NewIndexController()
			for n := 0; n < 10000; n++ {
				ctl.integrity()
			}
		})
		t.Run("mem", func(t *testing.T) {
			// go tool pprof --pdf ~/go/bin/yourbinary /var/path/to/cpu.pprof > file.pdf
			defer profile.Start(profile.MemProfile).Stop()
			ctl := NewIndexController()
			for n := 0; n < 10000; n++ {
				ctl.integrity()
			}
		})
	})
}

func BenchmarkIndexMethods(b *testing.B) {
	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = NewIndexController()
		}
	})
	b.Run("status", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		for n := 0; n < b.N; n++ {
			ctl.status()
		}
	})

	b.Run("status-reload", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		for n := 0; n < b.N; n++ {
			ctl.statusReload()
		}
	})

	b.Run("integrity", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		for n := 0; n < b.N; n++ {
			ctl.integrity()
		}
	})

	b.Run("integrity-reload", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		for n := 0; n < b.N; n++ {
			ctl.integrityReload()
		}
	})
}
