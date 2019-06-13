package auth

import (
	"errors"

	"github.com/zerjioang/etherniti/shared/mixed"

	"github.com/zerjioang/etherniti/core/modules/stack"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// new login request dto
type AuthRequest struct {
	mixed.DatabaseObjectInterface `json:"_,omitempty"`
	Uuid                          string             `json:"sid,omitempty"`
	Username                      string             `json:"name,omitempty" form:"name" query:"name"`
	Role                          constants.UserRole `json:"role,omitempty" form:"role" query:"role"`
	Email                         string             `json:"email" form:"email" query:"email"`
	Password                      string             `json:"pwd" form:"pwd" query:"pwd"`
}

// implementation of interface DatabaseObjectInterface
func (req *AuthRequest) Key() []byte {
	return str.UnsafeBytes(str.ToLowerAscii(req.Email))
}
func (req *AuthRequest) Value() []byte {
	return str.GetJsonBytes(req)
}
func (req *AuthRequest) New() mixed.DatabaseObjectInterface {
	return NewEmptyAuthRequestPtr()
}

// custom validation logic for read operation
// return nil if everyone can read
func (req *AuthRequest) CanRead(context *echo.Context, key string) error {
	return nil
}

// custom validation logic for update operation
// return nil if everyone can update
func (req *AuthRequest) CanUpdate(context *echo.Context, key string) error {
	if context.User().Role() != constants.AdminUser {
		return data.ErrNotAuthorized
	}
	return nil
}

// custom validation logic for delete operation
// return nil if everyone can delete
func (req *AuthRequest) CanDelete(context *echo.Context, key string) error {
	return nil
}

// custom validation logic for write operation
// return nil if everyone can write
func (req *AuthRequest) CanWrite(context *echo.Context) error {
	if req.Email != "" && req.Password != "" && req.Username != "" {
		logger.Info("registering user with email: ", req.Email)
		// hash user password
		req.Password = db.Hash(req.Password)
		return nil
	}
	return errors.New("you have to provide a valid email, password and username")
}

// custom validation logic for list operation
// return nil if everyone can list
func (req *AuthRequest) CanList(context *echo.Context) error {
	return data.ErrListingNotSupported
}

func (req *AuthRequest) Bind(context *echo.Context) (mixed.DatabaseObjectInterface, stack.Error) {
	if err := context.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return nil, stack.Ret(err)
	}
	return nil, data.ErrBind
}

func NewEmptyAuthRequestPtr() *AuthRequest {
	return new(AuthRequest)
}

func NewEmptyAuthRequest() AuthRequest {
	return AuthRequest{}
}

func NewDBAuthModel() mixed.DatabaseObjectInterface {
	return NewEmptyAuthRequestPtr()
}
