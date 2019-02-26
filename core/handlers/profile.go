// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"

	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/eth/counter"
	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/labstack/echo"
)

const (
	defaultProfileRequestTime = cache.DefaultExpiration
	readErr                   = `there was an trycatch during execution`
	bindErr                   = `there was an trycatch while processing your request information`
	itemDeleted               = `profile entry successfully deleted`
	noExistsNoUpdate          = `there was an trycatch during execution and could not update requeste profile`
	itemUpdated               = `profile entry successfully updated`
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
	req := api.NewProfileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding trycatch
		logger.Error("failed to bind request data to model: ", err)
		return ErrorStr(c, bindErr)
	}
	// create the connection profile
	userProfile := profile.NewConnectionProfileWithData(req)
	// create the token
	token, err := profile.CreateConnectionProfileToken(userProfile)
	if err == nil {
		rawBytes := util.GetJsonBytes(api.NewApiResponse("profile token successfully created", token))
		// increment created counter
		profilesCreated.Increment()
		return c.JSONBlob(http.StatusOK, rawBytes)
	} else {
		//token generation trycatch
		rawBytes := util.GetJsonBytes(api.NewApiError(http.StatusBadRequest, err.Error()))
		return c.JSONBlob(http.StatusOK, rawBytes)
	}
}

// profile validation check
func (ctl ProfileController) validate(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, util.Bytes(readErr))
}

// profile validation check
func (ctl ProfileController) count(c echo.Context) error {
	return c.JSON(
		http.StatusOK, profilesCreated.Get(),
	)
}

// new profile delete request
func (ctl ProfileController) clear(c echo.Context) error {
	// read target profile selection by user id
	//targetId := c.Param("id")
	// remove requested id from cache
	//ctl.cache.Delete(targetId)
	return c.JSONBlob(http.StatusOK, util.Bytes(itemDeleted))
}

// implemented method from interface RouterRegistrable
func (ctl ProfileController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing profile controller methods")
	router.POST("/profile", ctl.create)
	router.GET("/profile/count", ctl.count)
	router.GET("/profile/validate", ctl.validate)
}
