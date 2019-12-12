// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dashboard

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/auth"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type ApiKeysAuthController struct {
	common.DatabaseController
}

// constructor like function
func NewApiKeysAuthController() ApiKeysAuthController {
	uiCtl := ApiKeysAuthController{}
	var err error
	uiCtl.DatabaseController, err = common.NewDatabaseController("", "apikeys", auth.NewDBAuthModel)
	if err != nil {
		logger.Error("failed to create authentication controller ", err)
	}
	return uiCtl
}

// login with provided api key and secret pair
func (ctl ApiKeysAuthController) login(c *echo.Context) error {
	//new login request
	req := auth.NewEmptyAuthRequest()
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	if req.ApiKey != "" && req.ApiSecret != "" {
		logger.Info("logging user with api key and secret: ", req.ApiKey)
		item, readErr := ctl.GetKey([]byte(req.ApiKey))
		if readErr == nil {
			dto := auth.NewEmptyAuthRequest()
			pErr := db.Unserialize(item, &dto)
			if pErr != nil {
				logger.Error("failed to unserialize api key data: ", pErr.Error())
				return api.ErrorBytes(c, data.DatabaseError)
			}
			// check if provided api secret matches with stored one
			matches := req.ApiSecret == dto.ApiSecret
			if matches {
				// create authentication token
				token, err := createToken(dto.Uuid)
				if err != nil || token == "" {
					logger.Error("failed to create authentication token: ", err)
					return api.ErrorBytes(c, data.InvalidLoginData)
				}
				return api.SendSuccess(c, data.UserLogin, auth.NewLoginResponse(token))
			}
			return api.ErrorBytes(c, data.InvalidLoginData)
		}
		if readErr.Error() == "Key not found"{
			return api.ErrorBytes(c, data.InvalidLoginAPIKeyData)
		}
		// internal server erro while executing db call operation
		return api.ErrorBytesWithCode(c, protocol.StatusInternalServerError, data.InvalidLoginAPIKeyData)
	}
	logger.Error("failed to generate user token due to missing api key or secret in the request")
	return api.ErrorBytes(c, data.MissingAPIKeyLoginFields)
}

func (ctl ApiKeysAuthController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing api-key based authentication controller methods")
	router.POST("/auth/key", ctl.login)
}