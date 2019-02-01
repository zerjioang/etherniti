package profile

import (
	"github.com/zerjioang/gaethway/core/util"
	"testing"
	"time"
)

func TestCreateConnectionProfileToken(t *testing.T) {
	t.Run("empty-connection-profile", func(t *testing.T) {
		NewConnectionProfile()
	})
	t.Run("create-token", func(t *testing.T) {
		now := time.Now()
		p := ConnectionProfile{
			ConnectionId: util.GenerateUUID(),
			NodeAddress: "http://127.0.0.1:8454",
			Mode: "http",
			Port: 8454,
			Account: "0x0",
			//standard claims
			Id: util.GenerateUUID(),
			Issuer: "gaethway",
			ExpiresAt: now.Add(10*time.Minute).Unix(),
			NotBefore: now.Unix(),
			IssuedAt: now.Unix(),
		}
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
				t.Error("failed to control error for empty tokens")
			}
		})
		t.Run("parse-token", func(t *testing.T) {
			profile, err := ParseConnectionProfileToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb25uZWN0aW9uX2lkIjoiNjZkN2ZmZGYtZTRlNy00NTQ1LWJiMTYtY2QxNmJmNDM1ODUyIiwibm9kZV9hZGRyZXNzIjoiaHR0cDovLzEyNy4wLjAuMTo4NDU0IiwibW9kZSI6Imh0dHAiLCJwb3J0Ijo4NDU0LCJhY2NvdW50IjoiMHgwIiwiZXhwIjoxNTQ5MDIzNjM2LCJqdGkiOiI4ZDc5YTNhMC0xYzZhLTRjOGUtYTM4NS02N2M0NTlmMGM0ZTQiLCJpYXQiOjE1NDkwMjMwMzYsImlzcyI6ImdhZXRod2F5IiwibmJmIjoxNTQ5MDIzMDM2fQ.PXIQarefAxSCyClRhKpJisd3A0xZ-gPkVUgnGDPP474")
			if err != nil {
				t.Error(err)
			}
			t.Log(profile)
		})
	})
}
