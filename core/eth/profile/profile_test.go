package profile

import (
	"testing"
)

var (
	testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub2RlX2FkZHJlc3MiOiJodHRwOi8vMTI3LjAuMC4xOjg0NTQiLCJtb2RlIjoiaHR0cCIsInBvcnQiOjg0NTQsImFjY291bnQiOiIweDAiLCJleHAiOjEwNTUwMjUwMjA5LCJqdGkiOiJmNTI3N2M2Zi1jZDNmLTRkZTUtOTlkZC1hMTQ0YjgyNDJhYTEiLCJpYXQiOjE1NTAyNTAyMDksImlzcyI6ImV0aGVybml0aSIsIm5iZiI6MTU1MDI1MDIwOX0.MngYg9AI6ozwLBylild52vgpiwLYBqwjyVT7oSSRglg"
)

func TestCreateConnectionProfileToken(t *testing.T) {
	t.Run("empty-connection-profile", func(t *testing.T) {
		NewConnectionProfile()
	})
	t.Run("connection-profile", func(t *testing.T) {
		_ = NewDefaultConnectionProfile()
	})
	t.Run("create-token", func(t *testing.T) {
		p := NewDefaultConnectionProfile()
		token, err := CreateConnectionProfileToken(p)
		if err != nil {
			t.Error(err)
		}
		t.Log(token)
	})
	t.Run("parse-token", func(t *testing.T) {
		t.Run("parse-empty", func(t *testing.T) {
			_, err := ParseConnectionProfileToken("")
			if err == nil {
				t.Error("failed to control trycatch for empty tokens")
			}
		})
		t.Run("parse-token", func(t *testing.T) {
			profile, err := ParseConnectionProfileToken(testToken)
			if err != nil {
				t.Error(err)
			}
			t.Log(profile)
		})
	})
}
