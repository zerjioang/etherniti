package auth

import "github.com/zerjioang/etherniti/core/util/str"

// new login response dto
type LoginResponse struct {
	Token string `json:"token"`
}

func (res LoginResponse) Json() []byte {
	return str.GetJsonBytes(res)
}

func NewLoginResponse(tokenStr string) LoginResponse {
	return LoginResponse{Token: tokenStr}
}
