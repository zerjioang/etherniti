// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package id

import "testing"

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

func TestGenerateID(t *testing.T) {
	t.Run("id-entropy", func(t *testing.T) {
		value := GenerateIDString()
		t.Log("uuid value:", value.String())
	})
}
