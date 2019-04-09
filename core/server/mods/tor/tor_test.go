// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tor

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestBlockTorConnections(t *testing.T) {
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
			if got := BlockTorConnections(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlockTorConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}
