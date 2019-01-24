// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/methw/shared"
)

type ProfileController struct {
	shared.AutoRouteable
	cache *cache.Cache
}

func NewProfileController() ProfileController {
	ctl := ProfileController{}
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	return ctl
}

// new profile create request
func (ctl ProfileController) create(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

// new profile read request
func (ctl ProfileController) read(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

// new profile update request
func (ctl ProfileController) update(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

// new profile delete request
func (ctl ProfileController) delete(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

// new profile list request
func (ctl ProfileController) list(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

// implemented method from interface RouterRegistrable
func (ctl ProfileController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing profile controller methods")
	router.POST("/profile", ctl.create)
	router.GET("/profile/:id", ctl.read)
	router.PUT("/profile/:id", ctl.update)
	router.DELETE("/profile/:id", ctl.delete)
	router.GET("/", ctl.list)
}
