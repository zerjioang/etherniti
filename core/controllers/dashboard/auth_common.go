package dashboard

import (
	"time"

	"github.com/zerjioang/etherniti/util/banner"
	"github.com/zerjioang/go-hpc/lib/fastime"
	jwt "github.com/zerjioang/go-hpc/thirdparty/jwt-go"
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
