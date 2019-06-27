package auth

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

func NewLoginResponse(tokenStr string) LoginResponse {
	return LoginResponse{Token: tokenStr}
}
