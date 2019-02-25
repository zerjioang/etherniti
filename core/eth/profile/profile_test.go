package profile

import (
	"github.com/zerjioang/etherniti/core/eth/fastime"
	"testing"
	"time"

	"github.com/zerjioang/etherniti/core/util"
)

var (
	testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb25uZWN0aW9uX2lkIjoiMjUwZDVkNDEtYzNmNy00MTc3LTk4MjItNzI3YWRlNDQyZWNkIiwibm9kZV9hZGRyZXNzIjoiaHR0cDovLzEyNy4wLjAuMTo4NDU0IiwibW9kZSI6Imh0dHAiLCJwb3J0Ijo4NDU0LCJhY2NvdW50IjoiMHgwIiwiZXhwIjoxOTEwMjE4MDYxLCJqdGkiOiI0N2U5YjVhYi1hNDg4LTRmNWEtYjg0Ny0zMTEwODVhNDA2NzIiLCJpYXQiOjE1NTAyMTgwNjEsImlzcyI6ImV0aGVybml0aSIsIm5iZiI6MTU1MDIxODA2MX0.BE8yxe35eVtbWNF_pwWrh7-vHIRWTUDya9kQ8dLchr0"
)

func TestCreateConnectionProfileToken(t *testing.T) {

	t.Run("empty-connection-profile", func(t *testing.T) {
		NewConnectionProfile()
	})

	t.Run("connection-profile", func(t *testing.T) {
		now := fastime.Now()
		unixtime := now.Unix()
		_ = ConnectionProfile {
			NodeAddress:  "http://127.0.0.1:8454",
			Mode:         "http",
			Port:         8454,
			Account:      "0x0",
			//standard claims
			Id:        util.GenerateUUID(),
			Issuer:    "etherniti",
			ExpiresAt: now.Add(10 * time.Minute).Unix(),
			NotBefore: unixtime,
			IssuedAt:  unixtime,
		}
	})

	t.Run("create-token", func(t *testing.T) {
		now := fastime.Now()
		p := ConnectionProfile{
			NodeAddress:  "http://127.0.0.1:8454",
			Mode:         "http",
			Port:         8454,
			Account:      "0x0",
			//standard claims
			Id:        util.GenerateUUID(),
			Issuer:    "etherniti",
			ExpiresAt: now.Add(10 * time.Minute).Unix(),
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
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
			profile, err := ParseConnectionProfileToken(testToken)
			if err != nil {
				t.Error(err)
			}
			t.Log(profile)
		})
	})
}
