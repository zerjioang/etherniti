// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"encoding/base64"
	"errors"
	"sync/atomic"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/solc"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// token controller
type SolcController struct {
	//cached value. concurrent safe that stores []byte
	solidityVersionResponse atomic.Value
}

// constructor like function
func NewSolcController() SolcController {
	ctl := SolcController{}
	return ctl
}

func (ctl *SolcController) version(c *echo.Context) error {
	v := ctl.solidityVersionResponse.Load()
	if v == nil {
		// value not set. generate and store in cache
		// generate value
		solData, err := solc.SolidityVersion()
		if solData == nil {
			return errors.New("failed to get solc version")
		} else if err != nil {
			return err
		} else {
			// store in cache
			versionResponse := api.ToSuccess(data.SolcVersion, solData)
			// cache response for next request
			ctl.solidityVersionResponse.Store(versionResponse)
			// return response to client
			return api.SendSuccessBlob(c, versionResponse)
		}
	} else {
		//value already set and stored in memory cache
		versionResponse := v.([]byte)
		// return response to client
		return api.SendSuccessBlob(c, versionResponse)
	}
}

func (ctl SolcController) compileSingle(c *echo.Context) error {
	//read request parameters encoded in the body
	req := protocol.ContractCompileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}

	if req.Contract == "" {
		return errors.New("failed to get request contract data")
	} else {
		//compile given contract
		compilerResponse, err := solc.CompileSolidityString(req.Contract)
		if err != nil {
			return api.Error(c, err)
		} else {
			return api.SendSuccess(c, data.SolcCompiled, compilerResponse)
		}
	}
}

func (ctl SolcController) compileSingleFromBase64(c *echo.Context) error {
	//read request parameters encoded in the body
	req := protocol.ContractCompileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}

	if req.Contract == "" {
		return errors.New("failed to get request contract data")
	} else {
		//decode contract from base64 string to ascii
		decoded, b64Err := base64.StdEncoding.DecodeString(req.Contract)
		if b64Err != nil {
			//decoding error found
			return api.Error(c, b64Err)
		} else {
			//decoding success. compile the contract
			compilerResponse, err := solc.CompileSolidityString(string(decoded))
			if err != nil {
				return api.Error(c, err)
			} else {
				//generate the response for the client
				var contractResp []protocol.ContractCompileResponse
				contractResp = make([]protocol.ContractCompileResponse, len(compilerResponse))
				idx := 0
				for _, v := range compilerResponse {
					//populate current response
					current := contractResp[idx]
					current.Code = v.Code
					current.RuntimeCode = v.RuntimeCode
					current.Language = v.Info.Language
					current.LanguageVersion = v.Info.LanguageVersion
					current.CompilerVersion = v.Info.CompilerVersion
					current.CompilerOptions = v.Info.CompilerOptions
					current.AbiDefinition = v.Info.AbiDefinition
					contractResp[idx] = current
					idx++
				}
				return api.SendSuccess(c, data.SolcCompiled, contractResp)
			}
		}
	}
}

func (ctl SolcController) compileMultiple(c *echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, data.SolcVersion, solData)
	}
}

func (ctl SolcController) compileFromGit(c *echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, data.SolcVersion, solData)
	}
}

func (ctl SolcController) compileFromUploadedZip(c *echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return errors.New("failed to get solc version")
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, data.SolcVersion, solData)
	}
}

func (ctl SolcController) compileFromUploadedTargz(c *echo.Context) error {
	solData, err := solc.SolidityVersion()
	if solData == nil {
		return data.ErrCannotReadSolcVersion
	} else if err != nil {
		return err
	} else {
		return api.SendSuccess(c, data.SolcVersion, solData)
	}
}

func (ctl SolcController) compileModeSelector(c *echo.Context) error {
	mode := c.Param("mode")
	mode = str.ToLowerAscii(mode)
	logger.Debug("compiling ethereum contract with mode: ", mode)
	switch mode {
	case data.SingleRawFile:
		logger.Debug("compiling ethereum contract from raw solidity source code")
		return ctl.compileSingle(c)
	case data.SingleBase64File:
		logger.Debug("compiling ethereum contract from raw solidity encoded base64 code")
		return ctl.compileSingleFromBase64(c)
	case data.GitMode:
		logger.Debug("compiling ethereum contract from remote git repository")
		return ctl.compileFromGit(c)
	case data.ZipMode:
		logger.Debug("compiling ethereum contract from user provided zip file")
		return ctl.compileFromUploadedZip(c)
	case data.TargzMode:
		logger.Debug("compiling ethereum contract from user provided targz file")
		return ctl.compileFromUploadedTargz(c)
	default:
		return data.ErrUnknownMode
	}
}

// implemented method from interface RouterRegistrable
func (ctl SolcController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing solc controller methods")
	router.GET("/solc/version", ctl.version)
	router.POST("/solc/compile/:mode", ctl.compileModeSelector)
}
