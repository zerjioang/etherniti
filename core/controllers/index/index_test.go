// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package index

import (
	"testing"
	"time"

	"github.com/zerjioang/etherniti/shared"

	"github.com/pkg/profile"
	"github.com/zerjioang/etherniti/core/listener/common"
)

func TestIndexConcurrency(t *testing.T) {
	t.Run("index-single-echo", func(t *testing.T) {
		testServer := common.NewServer(nil)
		c := shared.AdquireContext(common.NewContext(testServer))
		err := NewIndexController().Index(c)
		if err != nil {
			t.Log(err)
		}
		shared.ReleaseContext(c)
	})
	t.Run("index-concurrency-echo", func(t *testing.T) {
		times := 100
		testServer := common.NewServer(nil)
		ctl := NewIndexController()
		for i := 0; i < times; i++ {
			go func() {
				c := shared.AdquireContext(common.NewContext(testServer))
				err := ctl.Index(c)
				if err != nil {
					t.Log(err)
				}
				shared.ReleaseContext(c)
			}()
		}
		time.Sleep(2 * time.Second)
	})

	t.Run("status-single-echo", func(t *testing.T) {
		testServer := common.NewServer(nil)
		ctl := NewIndexController()
		c := shared.AdquireContext(common.NewContext(testServer))
		err := ctl.Status(c)
		if err != nil {
			t.Log(err)
		}
		shared.ReleaseContext(c)
	})
	t.Run("status-concurrency-echo", func(t *testing.T) {
		times := 100
		testServer := common.NewServer(nil)
		ctl := NewIndexController()
		for i := 0; i < times; i++ {
			go func() {
				c := shared.AdquireContext(common.NewContext(testServer))
				err := ctl.Status(c)
				if err != nil {
					t.Log(err)
				}
				shared.ReleaseContext(c)
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
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewIndexController()
		}
	})
	b.Run("status", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ctl.status()
		}
	})

	b.Run("integrity", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ctl.integrity()
		}
	})
}
