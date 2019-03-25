// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/solc"
	"github.com/zerjioang/etherniti/core/util"
)

// token controller
type SolcController struct {
}

// constructor like function
func NewSolcController() SolcController {
	ctl := SolcController{}
	return ctl
}

func (ctl SolcController) version(c echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, "solc version", solData)
	}
}

func (ctl SolcController) compileSingle(c echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, "solc version", solData)
	}
}

func (ctl SolcController) compileMultiple(c echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, "solc version", solData)
	}
}

func (ctl SolcController) compileFromGit(c echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, "solc version", solData)
	}
}

func (ctl SolcController) compileFromUpload(c echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, "solc version", solData)
	}
}

func (ctl SolcController) compileModeSelector(c echo.Context) error {
	mode := c.Param("source")
	mode = util.ToLowerAscii(mode)
	logger.Debug("Compiling ethereum contract with mode: ", mode)
	switch mode {
	case "single":
		logger.Debug("Compiling ethereum contract with single mode")
	default:
		return errors.New("unknown mode selected. Allowed modes are: single, git, zip, targz")
	}
	return nil
}

// implemented method from interface RouterRegistrable
func (ctl SolcController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing solc controller methods")
	router.GET("/solc/version", ctl.version)
	router.POST("/solc/compile/:source", ctl.compileModeSelector)
}
