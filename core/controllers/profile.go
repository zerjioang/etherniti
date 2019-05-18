// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/modules/counter32"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type ProfileController struct {
}

var (
	//atomic counters stored on heap
	profilesCreated counter32.Count32
)

func NewProfileController() ProfileController {
	ctl := ProfileController{}
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	//ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	return ctl
}

// new profile create request
func (ctl ProfileController) create(c *echo.Context) error {
	//new profile request
	req := protocol.ProfileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	// add current user IP to request
	req.Ip = c.RealIP()
	// create the connection profile
	userProfile := profile.NewConnectionProfileWithData(req)
	// create the token
	token, err := profile.CreateConnectionProfileToken(userProfile)
	if err == nil {
		rawBytes := str.GetJsonBytes(protocol.NewApiResponse(data.ProfileTokenSuccess, token))
		// increment created counter
		profilesCreated.Increment()
		return c.JSONBlob(protocol.StatusOK, rawBytes)
	} else {
		//token generation stack
		rawBytes := str.GetJsonBytes(protocol.NewApiError(protocol.StatusBadRequest, str.UnsafeBytes(err.Error())))
		return c.JSONBlob(protocol.StatusOK, rawBytes)
	}
}

// profile validation check
func (ctl ProfileController) validate(c *echo.Context) error {
	return c.JSONBlob(protocol.StatusOK, data.ReadErr)
}

// profile validation counter
func (ctl ProfileController) count(c *echo.Context) error {
	c.OnSuccessCachePolicy = 10
	return api.SendSuccessBlob(c, profilesCreated.JsonBytes())
}

// implemented method from interface RouterRegistrable
func (ctl ProfileController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing profile controller methods")
	router.POST("/profile", ctl.create)
	router.GET("/profile/count", ctl.count)
	router.GET("/profile/validate", ctl.validate)
}
