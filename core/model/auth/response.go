package auth

import (
	"io"

	"github.com/zerjioang/etherniti/core/util/str"
)

const (
	beginStr = `{"token":"`
	endStr   = `"}`
)

var (
	begin = []byte(beginStr)
	end   = []byte(endStr)
)

// new login response dto
type LoginResponse struct {
	Token string `json:"token"`
}

func (res LoginResponse) Json() []byte {
	return str.UnsafeBytes(beginStr + res.Token + endStr)
}

func (res LoginResponse) Writer(w io.Writer) error {
	if w != nil {
		_, _ = w.Write(begin)
		_, _ = w.Write(str.UnsafeBytes(res.Token))
		_, _ = w.Write(end)
	}
	return nil
}

func NewLoginResponse(tokenStr string) LoginResponse {
	return LoginResponse{Token: tokenStr}
}
