// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profile

import (
	"errors"

	"github.com/zerjioang/etherniti/core/util/ip"
	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/util/id"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/thirdparty/jwt-go"
)

var (
	tokenSecretBytes        = []byte(config.TokenSecret())
	errTokenNoValid         = errors.New("provided token contains invalid or missing fields")
	errInvalidSigningMethod = errors.New("unexpected signing method")
	errFailedToRead         = errors.New("failed to read token claims")
)

// default data for connection profile
type ConnectionProfile struct {
	jwt.Claims `json:"_,omitempty"`
	protocol.ProfileRequest

	// service version when profile was generated
	Version string `json:"version"`
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
	// validity of the profile: whether all required data is present or not
	Valididity bool `json:"validity"`
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
	profile.RpcEndpoint = profile.readString(claims["endpoint"])
	profile.Address = profile.readString(claims["address"])
	profile.Key = profile.readString(claims["key"])
	profile.Version = constants.Version
	profile.Valididity = profile.readBool(claims["validity"])
	profile.Audience = profile.readString(claims["aud"])
	profile.ExpiresAt = profile.readInt64(claims["exp"])
	profile.Id = profile.readString(claims["jti"])
	profile.IssuedAt = profile.readInt64(claims["iat"])
	profile.Issuer = profile.readString(claims["iss"])
	profile.NotBefore = profile.readInt64(claims["nbf"])
	profile.Subject = profile.readString(claims["sub"])
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

func (profile ConnectionProfile) readUint16(v interface{}) uint16 {
	if v != nil {
		val, ok := v.(uint16)
		if ok {
			return val
		}
	}
	return 0
}

func (profile ConnectionProfile) readInt64(v interface{}) int64 {
	if v != nil {
		val, ok := v.(int64)
		if ok {
			return val
		}
	}
	return 0
}

func (profile ConnectionProfile) readBool(v interface{}) bool {
	if v != nil {
		val, ok := v.(bool)
		if ok {
			return val
		}
	}
	return false
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
			return nil, errInvalidSigningMethod
		}
		// return used token secret
		return tokenSecretBytes, nil
	})
	if err != nil {
		return profile, err
	}

	mapc, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return profile, errFailedToRead
	}
	profile = profile.Populate(mapc)

	//check profile validity
	profile.Valididity = profile.RpcEndpoint != ""

	if profile.Valididity {
		return profile, nil
	} else {
		return profile, errTokenNoValid
	}
}

//constructor like function
func NewConnectionProfile() ConnectionProfile {
	p := ConnectionProfile{}
	return p
}

//constructor like function
func NewConnectionProfileWithData(data protocol.ProfileRequest) ConnectionProfile {
	now := fastime.Now()
	p := ConnectionProfile{
		Id: id.GenerateUUIDFromEntropy(),
		ProfileRequest: protocol.ProfileRequest{
			RpcEndpoint: data.RpcEndpoint,
			Address:     data.Address, //required
			Key:         data.Key,
			Source:      ip.Ip2int(data.Ip),
		},
		Issuer:    "etherniti.org",
		ExpiresAt: now.Add(config.TokenExpiration()).Unix(),
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
		Version:   constants.Version,
	}
	//check profile validity
	p.Valididity = p.Id != "" &&
		p.RpcEndpoint != "" &&
		p.Address != "" &&
		p.Key != ""
	return p
}

func NewDefaultConnectionProfile() ConnectionProfile {
	now := fastime.Now()
	return ConnectionProfile{
		ProfileRequest: protocol.ProfileRequest{
			RpcEndpoint: "http://127.0.0.1:8545",
			Address:     "0x0",
			Key:         "0x0",
		},
		//standard claims
		Id:         id.GenerateUUIDFromEntropy(),
		Issuer:     "etherniti",
		ExpiresAt:  now.Add(10 * fastime.Minute).Unix(),
		NotBefore:  now.Unix(),
		IssuedAt:   now.Unix(),
		Valididity: false,
	}
}
