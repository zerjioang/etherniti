package ui

import (
	"errors"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// new login request dto
type AuthRequest struct {
	common.DatabaseObjectInterface
	Username string             `json:"user,omitempty" form:"user" query:"user"`
	Role     constants.UserRole `json:"user,omitempty" form:"user" query:"user"`
	Email    string             `json:"email" form:"email" query:"email"`
	Password string             `json:"pwd" form:"pwd" query:"pwd"`
}

// implementation of interface DatabaseObjectInterface
func (req *AuthRequest) Key() []byte {
	return str.UnsafeBytes(str.ToLowerAscii(req.Email))
}
func (req *AuthRequest) Value() []byte {
	return str.GetJsonBytes(req)
}
func (req *AuthRequest) New() common.DatabaseObjectInterface {
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
	return data.ListingNotSupported
}

func (req *AuthRequest) Bind(context *echo.Context) common.DatabaseObjectInterface {
	if err := context.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		_ = api.ErrorStr(context, data.BindErr)
	}
	return nil
}

func NewEmptyAuthRequestPtr() *AuthRequest {
	return new(AuthRequest)
}

func NewEmptyAuthRequest() AuthRequest {
	return AuthRequest{}
}

func NewDBAuthModel() common.DatabaseObjectInterface {
	return NewEmptyAuthRequestPtr()
}
