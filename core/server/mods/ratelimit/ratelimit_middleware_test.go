// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ratelimit

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestRateLimit(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RateLimit(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RateLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
