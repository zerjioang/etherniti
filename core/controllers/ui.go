// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/controllers/ui"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type UIAuthController struct {
	common.DatabaseController
}

// constructor like function
func NewUIAuthController() UIAuthController {
	uiCtl := UIAuthController{}
	var err error
	uiCtl.DatabaseController, err = common.NewDatabaseController("auth", ui.NewDBAuthModel())
	if err != nil {
		logger.Error("failed to create database controller ", err)
	}
	return uiCtl
}

// logins user data and returns access token
func (ctl UIAuthController) login(c *echo.Context) error {

	//new login request
	req := ui.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" {
		logger.Info("logging user with email: ", req.Email)
		item, readErr := ctl.GetKey([]byte(req.Email))
		if readErr == nil {
			dto := ui.NewEmptyAuthRequest()
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
func (ctl UIAuthController) register(c *echo.Context) error {
	//new login request
	req := ui.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	if req.Email != "" && req.Password != "" && req.Username != "" {
		logger.Info("registering user with email: ", req.Email)
		// hash user password
		req.Password = db.Hash(req.Password)
		saveErr := ctl.SetUniqueKey(str.UnsafeBytes(req.Email), db.Serialize(req))
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
func (ctl UIAuthController) recaptcha(c *echo.Context) error {
	return data.ErrNotImplemented
}

func (ctl UIAuthController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing ui controller methods")
	router.POST("/ui/login", ctl.login)
	router.POST("/ui/register", ctl.register)
	router.POST("/ui/recaptcha", ctl.recaptcha)
}
