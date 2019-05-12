// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type UIController struct {
}

// contructor like function
func NewUIController() UIController {
	dc := UIController{}
	return dc
}

// logins user data and returns access token
func (ctl UIController) login(c *echo.Context) error {

	//new login request
	req := protocol.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" {
		logger.Info("logging user with email: ", req.Email)
		item, readErr := db.GetInstance().GetKeyValue([]byte(req.Email))
		if readErr == nil {
			dto := new(protocol.RegisterRequest)
			pErr := db.Unserialize(item, dto)
			if pErr != nil {
				return api.ErrorStr(c, data.DatabaseError)
			} else {
				// check if email and password matches
				matches := req.Email == dto.Email && db.CompareHash(req.Password, dto.Password)
				if matches {
					return api.Success(c, data.UserLogin, str.UnsafeBytes("token"))
				} else {
					return api.ErrorStr(c, data.InvalidLoginData)
				}
			}
		} else {
			//db read error
			// this code is trigger each time user fails a login attempt
			return api.ErrorStr(c, data.FailedLoginVerification)
		}
	} else {
		return api.ErrorStr(c, data.MissingLoginFields)
	}
}

// registers new user data in the api sever
func (ctl UIController) register(c *echo.Context) error {
	//new login request
	req := protocol.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" && req.Username != "" {
		logger.Info("registering user with email: ", req.Email)
		// hash user password
		req.Password = db.Hash(req.Password)
		saveErr := db.GetInstance().PutUniqueKeyValue(str.UnsafeBytes(req.Email), db.Serialize(req))
		if saveErr != nil {
			logger.Error("failed to register new user due to: ", saveErr)
			return api.ErrorStr(c, data.UserRegisterFailed)
		} else {
			return api.Success(c, data.UserRegistered, data.RegistrationSuccess)
		}
	}
	return nil
}

// validates recatpcha requests
// more info at: https://www.google.com/recaptcha/admin/site/346227166
func (ctl UIController) recaptcha(c *echo.Context) error {
	return data.ErrNotImplemented
}

func (ctl UIController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing ui controller methods")
	router.POST("/ui/login", ctl.login)
	router.POST("/ui/register", ctl.register)
	router.POST("/ui/recaptcha", ctl.recaptcha)
}
