// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package listener

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

func TestFactoryListener(t *testing.T) {
	type args struct {
		typeof listener.ServiceType
	}
	tests := []struct {
		name string
		args args
		want listener.ListenerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactoryListener(tt.args.typeof); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FactoryListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
