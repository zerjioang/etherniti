// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/handlers/clientcache"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/etherniti/core/eth/counter"
	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/logger"
)

const (
	readErr = `there was an error during execution`
	bindErr = `there was an error while processing your request information`
)

type ProfileController struct {
	//cache *cache.Cache
}

var (
	//atomic counters stored on heap
	profilesCreated counter.Count32
)

func NewProfileController() ProfileController {
	ctl := ProfileController{}
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	//ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	return ctl
}

// new profile create request
func (ctl ProfileController) create(c echo.Context) error {
	//new profile request
	req := protocol.ProfileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, bindErr)
	}
	// add current user IP to request
	req.Ip = c.RealIP()
	// create the connection profile
	userProfile := profile.NewConnectionProfileWithData(req)
	// create the token
	token, err := profile.CreateConnectionProfileToken(userProfile)
	if err == nil {
		rawBytes := str.GetJsonBytes(protocol.NewApiResponse("profile token successfully created", token))
		// increment created counter
		profilesCreated.Increment()
		return c.JSONBlob(http.StatusOK, rawBytes)
	} else {
		//token generation trycatch
		rawBytes := str.GetJsonBytes(protocol.NewApiError(http.StatusBadRequest, err.Error()))
		return c.JSONBlob(http.StatusOK, rawBytes)
	}
}

// profile validation check
func (ctl ProfileController) validate(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, str.UnsafeBytes(readErr))
}

// profile validation check
func (ctl ProfileController) count(c echo.Context) error {
	var code int
	code, c = clientcache.Cached(c, true, 10) // cache policy: 10 seconds
	return c.JSON(code, profilesCreated.Get())
}

// implemented method from interface RouterRegistrable
func (ctl ProfileController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing profile controller methods")
	router.POST("/profile", ctl.create)
	router.GET("/profile/count", ctl.count)
	router.GET("/profile/validate", ctl.validate)
}
