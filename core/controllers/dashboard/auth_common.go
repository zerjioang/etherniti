package dashboard

import (
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/util/banner"
	"github.com/zerjioang/etherniti/thirdparty/jwt-go"
	"time"
)

func createToken(userUuid string) (string, error) {
	type Claims struct {
		User    string `json:"sid"`
		Version string `json:"version"`
		jwt.StandardClaims
	}
	// Declare the expiration time of the token
	now := fastime.Now()
	// here, we have kept it as 20 minutes
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		User: userUuid,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "etherniti.org",
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
		},
		Version: banner.Version,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(authTokenSecret)
}
