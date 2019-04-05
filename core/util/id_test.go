// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	t.Run("uuid-entropy", func(t *testing.T) {
		value := GenerateUUIDFromEntropy()
		t.Log("uuid value:", value)
		// example: 07dc4b26-caef-43c9-b068-54fff6222653
		if value == "" || len(value) != 36 {
			t.Error("failed to create uuid from entropy")
		}
	})
}

func TestGenerateUUIDFromEntropy(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUUIDFromEntropy(); got != tt.want {
				t.Errorf("GenerateUUIDFromEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}
