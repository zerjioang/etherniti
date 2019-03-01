// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profile

import (
	"errors"
	"fmt"

	"github.com/zerjioang/etherniti/core/api/protocol"

	"github.com/etherniti/jwt-go"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/core/util"
)

var (
	emptyProfile     ConnectionProfile
	tokenSecretBytes = []byte(config.TokenSecret)
	errTokenNoValid  = errors.New("invalid fields token provided")
)

// default data for connection profile
type ConnectionProfile struct {
	jwt.Claims `json:"_,omitempty"`

	//network id of target connection
	NetworkId uint8 `json:"networkId"`

	// address of the connection node: ip, domain, infura, etc
	Peer string `json:"peer"`

	//connection mode: ipc,http,rpc
	Mode string `json:"mode"`

	//connection por if required
	Port int `json:"port"`

	// default ethereum account for transactioning
	Address string `json:"address"`

	// user or device private key
	Key string `json:"key"`

	// service version when profile was generated
	Version int `json:"version"`

	// validity of the profile: whether all required data is present or not
	Valididity bool `json:"validity"`

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

// implementation from Claims
func (profile ConnectionProfile) Valid() error {
	if !profile.Valididity {
		return errTokenNoValid
	}
	return nil
}

func (profile ConnectionProfile) Secret() []byte {
	return tokenSecretBytes
}
func (profile ConnectionProfile) Populate(claims jwt.MapClaims) ConnectionProfile {
	profile.Peer = profile.readString(claims["peer"])
	profile.Mode = profile.readString(claims["mode"])
	profile.Address = profile.readString(claims["address"])
	profile.Key = profile.readString(claims["key"])
	profile.Version = profile.readInt(claims["version"])
	profile.Id = profile.readString(claims["id"])
	profile.Key = profile.readString(claims["key"])
	return profile
}

func (profile ConnectionProfile) readString(v interface{}) string {
	if v != nil {
		str, ok := v.(string)
		if ok {
			return str
		}
	}
	return ""
}

func (profile ConnectionProfile) readInt(v interface{}) int {
	if v != nil {
		val, ok := v.(int)
		if ok {
			return val
		}
	}
	return 0
}

func CreateConnectionProfileToken(profile ConnectionProfile) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, profile)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(tokenSecretBytes)
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

	mapc, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return profile, errors.New("failed to read token claims")
	}
	profile = profile.Populate(mapc)

	//check profile validity
	profile.Valididity = profile.Peer != "" &&
		profile.Address != "" &&
		profile.Key != ""

	if profile.Valididity {
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

//constructor like function
func NewConnectionProfileWithData(data protocol.NewProfileRequest) ConnectionProfile {
	now := fastime.Now()
	p := ConnectionProfile{
		Id:        util.GenerateUUID(),
		NetworkId: data.NetworkId,
		Peer:      data.Peer,    //required
		Address:   data.Address, // required
		Key:       data.Key,     //required
		Mode:      data.Mode,    //required
		Port:      data.Port,
		Issuer:    "etherniti.org",
		ExpiresAt: now.Add(config.TokenExpiration).Unix(),
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
		Version:   1,
	}
	//check profile validity
	p.Valididity = p.Id != "" &&
		p.Peer != "" &&
		p.Address != "" &&
		p.Key != ""
	return p
}

func NewDefaultConnectionProfile() ConnectionProfile {
	now := fastime.Now()
	return ConnectionProfile{
		Peer:    "http://127.0.0.1:8454",
		Mode:    "http",
		Port:    8454,
		Address: "0x0",
		Key:     "0x0",
		//standard claims
		Id:         util.GenerateUUID(),
		Issuer:     "etherniti",
		ExpiresAt:  now.Add(10 * fastime.Minute).Unix(),
		NotBefore:  now.Unix(),
		IssuedAt:   now.Unix(),
		Valididity: false,
	}
}
