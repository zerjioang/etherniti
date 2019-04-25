// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/helix/impl/util/str"
)

type UIController struct {
}

// contructor like function
func NewUIController() UIController {
	dc := UIController{}
	return dc
}

// logins user data and returns access token
func (ctl UIController) login(c echo.ContextInterface) error {

	//new login request
	req := protocol.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, constants.BindErr)
	}
	if req.Email != "" && req.Password != "" {
		logger.Info("logging user with email: ", req.Email)
		item, readErr := db.GetInstance().GetKeyValue([]byte(req.Email))
		if readErr == nil {
			dto := new(protocol.RegisterRequest)
			pErr := db.Unserialize(item, dto)
			if pErr != nil {
				return api.ErrorStr(c, "Failed to process your login request at this moment. Please try it later")
			} else {
				// check if email and password matches
				matches := req.Email == dto.Email && db.CompareHash(req.Password, dto.Password)
				if matches {
					return api.Success(c, "login", "token")
				} else {
					return api.ErrorStr(c, "Invalid username or password provided")
				}
			}
		} else {
			//db read error
			// this code is trigger each time user fails a login attempt
			return api.ErrorStr(c, "Failed to verify your login information at this time. Please try it few minutes later.")
		}
	} else {
		return api.ErrorStr(c, "Invalid login data provided. please fill all required fields")
	}
}

// registers new user data in the api sever
func (ctl UIController) register(c echo.ContextInterface) error {
	//new login request
	req := protocol.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, constants.BindErr)
	}
	if req.Email != "" && req.Password != "" && req.Username != "" {
		logger.Info("registering user with email: ", req.Email)
		// hash user password
		req.Password = db.Hash(req.Password)
		saveErr := db.GetInstance().PutUniqueKeyValue(str.ToBytes(req.Email), db.Serialize(req))
		if saveErr != nil {
			logger.Error("failed to register new user due to: ", saveErr)
			return api.ErrorStr(c, "failed to register new user account")
		} else {
			return api.Success(c, "user registered", req.Email)
		}
	}
	return nil
}

func (ctl UIController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing ui controller methods")
	router.POST("/ui/login", ctl.login)
	router.POST("/ui/register", ctl.register)
}
