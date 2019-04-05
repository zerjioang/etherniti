// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ratelimit

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/core/modules/cache"
)

func TestNewRateLimitEngine(t *testing.T) {
	tests := []struct {
		name string
		want RateLimitEngine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRateLimitEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRateLimitEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRateLimitEngine_Eval(t *testing.T) {
	type fields struct {
		rateCache *cache.MemoryCache
	}
	type args struct {
		clientIdentifier string
		h                http.Header
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   RateLimitResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rte := RateLimitEngine{
				rateCache: tt.fields.rateCache,
			}
			if got := rte.Eval(tt.args.clientIdentifier, tt.args.h); got != tt.want {
				t.Errorf("RateLimitEngine.Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
