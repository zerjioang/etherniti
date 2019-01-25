// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/methw/core/api"
	"github.com/zerjioang/methw/core/util"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type EthController struct {
	cache *cache.Cache
}

func NewEthController() EthController {
	ctl := EthController{}
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	return ctl
}

// new profile create request
func (ctl EthController) create(c echo.Context) error {
	//new profile request
	req := api.NewProfileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		return c.JSONBlob(http.StatusBadRequest, util.Bytes(bindErr))
	}
	// assign an unique uuid
	// save the data in the cache
	uuid := util.GenerateUUID()
	ctl.cache.Set(uuid, req, defaultProfileRequestTime)
	rawBytes := util.GetJsonBytes(api.NewApiResponse("profile entry successfully created", uuid))
	return c.JSONBlob(http.StatusOK, rawBytes)
}

// new profile read request
func (ctl EthController) read(c echo.Context) error {
	//new read profile request
	targetId := c.Param("id")
	//read the cache
	rawInterface, exists := ctl.cache.Get(targetId)
	if exists && rawInterface != nil {
		//serialize to json and return back
		rawBytes := util.GetJsonBytes(api.NewApiResponse("readed", rawInterface))
		return c.JSONBlob(http.StatusOK, rawBytes)
	}
	return c.JSONBlob(http.StatusOK, util.Bytes(readErr))
}

// new profile update request
func (ctl EthController) update(c echo.Context) error {
	//new profile request
	req := api.NewProfileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		return c.JSONBlob(http.StatusBadRequest, util.Bytes(bindErr))
	}
	// read target profile selection by user id
	targetId := c.Param("id")

	newProfile := api.Profile{}
	newProfile.Address = req.Address
	newProfile.Node = req.Node

	updateErr := ctl.cache.Replace(targetId, newProfile, defaultProfileRequestTime)
	if updateErr != nil {
		// return error
		return c.JSONBlob(http.StatusBadRequest, util.Bytes(noExistsNoUpdate))
	} else {
		// no update error thrown
		return c.JSONBlob(http.StatusOK, util.Bytes(itemUpdated))
	}
}

// new profile delete request
func (ctl EthController) delete(c echo.Context) error {
	// read target profile selection by user id
	targetId := c.Param("id")
	// remove requested id from cache
	ctl.cache.Delete(targetId)
	return c.JSONBlob(http.StatusOK, util.Bytes(itemDeleted))
}

// new profile list request
func (ctl EthController) list(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

// implemented method from interface RouterRegistrable
func (ctl EthController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing profile controller methods")
	router.POST("/profile", ctl.create)
	router.GET("/profile/:id", ctl.read)
	router.PUT("/profile/:id", ctl.update)
	router.DELETE("/profile/:id", ctl.delete)
	router.GET("/", ctl.list)
}
