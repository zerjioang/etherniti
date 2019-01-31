// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package profile

import (
	"fmt"
	"github.com/zerjioang/gaethway/core/keystore/memory"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type ConnectionProfile struct {
	ConnectionId string
	wallet       memory.WalletContent
}

func (profile ConnectionProfile) Claims() jwt.Claims {
	return jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}
}

func (profile ConnectionProfile) Secret() string {
	return "secret"
}

func NewConnectionProfile(claims jwt.MapClaims) (ConnectionProfile, error) {
	var profile ConnectionProfile
	if claims == nil {
		return profile, errors.New("failed to create connection profile with given token")
	}
	return profile, nil
}

func CreateConnectionProfileToken(profile ConnectionProfile) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, profile.Claims())

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(profile.Secret())
}

func ParseConnectionProfileToken(tokenStr string) (*ConnectionProfile, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return "", nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return NewConnectionProfile(claims)
	} else {
		return nil, err
	}
}
