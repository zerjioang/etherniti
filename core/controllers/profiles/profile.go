// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profiles

import (
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/etherniti/shared/dto"
	"github.com/zerjioang/etherniti/shared/notifier"
	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/lib/counter32"

	"github.com/zerjioang/etherniti/core/api"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

type ProfileController struct {
}

var (
	//atomic counters stored on heap
	// which supports concurrent access
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
func (ctl ProfileController) create(c *shared.EthernitiContext) error {
	//new profile request
	req := dto.ProfileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	// add current user IP to request
	req.Ip = c.RealIP()
	// create the connection profile
	userProfile := profile.NewConnectionProfileWithData(req)
	// create the token
	token, err := profile.CreateConnectionProfileToken(userProfile)
	if err == nil {
		// increment created counter
		profilesCreated.Increment()
		// send internal event: new profile created
		notifier.NewProfileRequestEvent.Emit()
		return api.SendSuccess(c, data.ProfileTokenSuccess, token)
	} else {
		//token generation error
		return api.Error(c, err)
	}
}

// profile validation check
func (ctl ProfileController) validate(c *shared.EthernitiContext) error {
	return c.JSONBlob(codes.StatusOK, data.ReadErr)
}

// profile validation counter
func (ctl ProfileController) count(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = 10
	return api.SendSuccessBlob(c, profilesCreated.JsonBytes())
}

// implemented method from interface RouterRegistrable
func (ctl ProfileController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing profile controller methods")
	router.POST("/profile", wrap.Call(ctl.create))
	router.GET("/profile/count", wrap.Call(ctl.count))
	router.GET("/profile/validate", wrap.Call(ctl.validate))
}
