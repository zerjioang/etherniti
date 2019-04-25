// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"
	"time"

	"github.com/pkg/profile"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/core/modules/concurrentbuffer"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestIndexConcurrency(t *testing.T) {
	t.Run("index-single-echo", func(t *testing.T) {
		testServer := common.NewServer(nil)
		ctx := common.NewContext(testServer)
		err := Index(ctx)
		if err != nil {
			t.Log(err)
		}
	})
	t.Run("index-concurrency-echo", func(t *testing.T) {
		times := 100
		testServer := common.NewServer(nil)
		for i := 0; i < times; i++ {
			go func() {
				ctx := common.NewContext(testServer)
				err := Index(ctx)
				if err != nil {
					t.Log(err)
				}
			}()
		}
		time.Sleep(2 * time.Second)
	})

	t.Run("status-single-echo", func(t *testing.T) {
		testServer := common.NewServer(nil)
		ctl := NewIndexController()
		ctx := common.NewContext(testServer)
		err := ctl.Status(ctx)
		if err != nil {
			t.Log(err)
		}
	})
	t.Run("status-concurrency-echo", func(t *testing.T) {
		times := 100
		testServer := common.NewServer(nil)
		ctl := NewIndexController()
		for i := 0; i < times; i++ {
			go func() {
				ctx := common.NewContext(testServer)
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

	b.Run("status-reload", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ctl.refreshStatusData()
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

	b.Run("integrity-reload", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		ctl := NewIndexController()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ctl.refreshIntegrityData()
		}
	})
}

func TestNewIndexController(t *testing.T) {
	tests := []struct {
		name string
		want *IndexController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		c echo.ContextInterface
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Index(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIndexController_Status(t *testing.T) {
	type fields struct {
		statusData    concurrentbuffer.ConcurrentBuffer
		integrityData concurrentbuffer.ConcurrentBuffer
	}
	type args struct {
		c echo.ContextInterface
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &IndexController{
				statusData:    tt.fields.statusData,
				integrityData: tt.fields.integrityData,
			}
			if err := ctl.Status(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("IndexController.Status() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIndexController_status(t *testing.T) {
	type fields struct {
		statusData    concurrentbuffer.ConcurrentBuffer
		integrityData concurrentbuffer.ConcurrentBuffer
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &IndexController{
				statusData:    tt.fields.statusData,
				integrityData: tt.fields.integrityData,
			}
			if got := ctl.status(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexController.status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexController_refreshStatusData(t *testing.T) {
	type fields struct {
		statusData    concurrentbuffer.ConcurrentBuffer
		integrityData concurrentbuffer.ConcurrentBuffer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &IndexController{
				statusData:    tt.fields.statusData,
				integrityData: tt.fields.integrityData,
			}
			ctl.refreshStatusData()
		})
	}
}

func TestIndexController_Integrity(t *testing.T) {
	type fields struct {
		statusData    concurrentbuffer.ConcurrentBuffer
		integrityData concurrentbuffer.ConcurrentBuffer
	}
	type args struct {
		c echo.ContextInterface
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &IndexController{
				statusData:    tt.fields.statusData,
				integrityData: tt.fields.integrityData,
			}
			if err := ctl.Integrity(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("IndexController.Integrity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIndexController_refreshIntegrityData(t *testing.T) {
	type fields struct {
		statusData    concurrentbuffer.ConcurrentBuffer
		integrityData concurrentbuffer.ConcurrentBuffer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &IndexController{
				statusData:    tt.fields.statusData,
				integrityData: tt.fields.integrityData,
			}
			ctl.refreshIntegrityData()
		})
	}
}

func TestIndexController_integrity(t *testing.T) {
	type fields struct {
		statusData    concurrentbuffer.ConcurrentBuffer
		integrityData concurrentbuffer.ConcurrentBuffer
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &IndexController{
				statusData:    tt.fields.statusData,
				integrityData: tt.fields.integrityData,
			}
			if got := ctl.integrity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexController.integrity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexController_RegisterRouters(t *testing.T) {
	type fields struct {
		statusData    concurrentbuffer.ConcurrentBuffer
		integrityData concurrentbuffer.ConcurrentBuffer
	}
	type args struct {
		router *echo.Group
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
			ctl := &IndexController{
				statusData:    tt.fields.statusData,
				integrityData: tt.fields.integrityData,
			}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
