// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profile

import (
	"errors"
	"fmt"
	"github.com/etherniti/jwt-go"
	"github.com/zerjioang/etherniti/core/config"
)

var (
	emptyProfile     ConnectionProfile
	defaultSigningMethod = jwt.SigningMethodHS256
	tokenSecretBytes = []byte(config.TokenSecret)
	errTokenNoValid  = errors.New("invalid fields token provided")
)

// default data for connection profile
type ConnectionProfile struct {
	jwt.Claims `json:"_,omitempty"`

	// address of the connection node: ip, domain, infura, etc
	NodeAddress string `json:"node_address"`

	//connection mode: ipc,http,rpc
	Mode string `json:"mode"`

	//connection por if required
	Port int `json:"port"`

	// default ethereum account for transactioning
	Account string `json:"account"`

	//standard claims
	//Identifies the recipients that the JWT is intended for.
	// Each principal intended to process the JWT must identify
	// itself with a value in the audience claim.
	// If the principal processing the claim does not identify
	// itself with a value in the aud claim when this claim is present,
	// then the JWT must be rejected.
	Audience string `json:"aud,omitempty"`
	// Identifies the expiration time on and after which the
	// JWT must not be accepted for processing. The value must be
	// a NumericDate[10]: either an integer or decimal, representing
	// seconds past 1970-01-01 00:00:00Z.
	ExpiresAt int64 `json:"exp,omitempty"`
	// Case sensitive unique identifier of the token
	// even among different issuers.
	// it works also as unique identifier of the connection profile
	Id string `json:"jti,omitempty"`
	// Identifies the time at which the JWT was issued.
	// The value must be a NumericDate.
	IssuedAt int64 `json:"iat,omitempty"`
	//Identifies principal that issued the JWT.
	Issuer string `json:"iss,omitempty"`
	// Identifies the time on which the JWT will start to be accepted
	// for processing. The value must be a NumericDate.
	NotBefore int64 `json:"nbf,omitempty"`
	//Identifies the subject of the JWT.
	Subject string `json:"sub,omitempty"`
}

func (profile ConnectionProfile) Valid() error {
	valid := profile.Id != "" &&
		profile.NodeAddress != "" &&
		profile.Account != ""
	if !valid {
		return errTokenNoValid
	}
	return nil
}

func (profile ConnectionProfile) Secret() []byte {
	return tokenSecretBytes
}

func CreateConnectionProfileToken(profile ConnectionProfile) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewTokenPoolWithClaims(defaultSigningMethod, profile)

	// Sign and get the complete encoded token as a string using the secret
	signedStr, err := token.SignedString(tokenSecretBytes)
	// return result
	return signedStr, err
}

func ParseConnectionProfileToken(tokenStr string) (ConnectionProfile, error) {
	var profile ConnectionProfile
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return used token secret
		return tokenSecretBytes, nil
	})
	if err != nil {
		return profile, err
	}

	profile, ok := token.Claims.(ConnectionProfile)
	if ok && token.Valid {
		return profile, nil
	} else {
		return profile, err
	}
}

//constructor like function
func NewConnectionProfile() ConnectionProfile {
	p := ConnectionProfile{}
	return p
}
