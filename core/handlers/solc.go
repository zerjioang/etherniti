// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"github.com/zerjioang/etherniti/core/modules/packers"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"gopkg.in/src-d/go-billy.v4/memfs"

	"sync/atomic"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/solc"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// token controller
type SolcController struct {
	//cached value. concurrent safe that stores []byte
	solidityVersionResponse atomic.Value
}
type CodeReader func(c *echo.Context) ([]string, error)

var (
	solcResponseCode    int
	solcVersionResponse []byte
	errNoContractData   = errors.New("failed to get request contract data")
	noContractData      []string
)

func init() {
	//preload solc version information
	solData, err := solc.SolidityVersion()
	if solData == nil {
		solcVersionResponse = api.ErrorBytes("failed to get solc version")
	} else if err != nil {
		solcResponseCode = protocol.StatusBadRequest
		solcVersionResponse = api.ErrorBytes(err.Error())
	} else {
		solcResponseCode = protocol.StatusOK
		solcVersionResponse = api.ToSuccess(data.SolcVersion, solData)
	}
}

// constructor like function
func NewSolcController() SolcController {
	ctl := SolcController{}
	return ctl
}

func (ctl *SolcController) version(c *echo.Context) error {
	return c.FastBlob(solcResponseCode, echo.MIMEApplicationJSONCharsetUTF8, solcVersionResponse)
}

// solc --optimize --optimize-runs 200 --opcodes --bin --abi --hashes --asm erc20.sol
func (ctl SolcController) compileFromSources(c *echo.Context, codeReader CodeReader) error {
	//read request parameters encoded in the body
	req := protocol.MultiFileContractCompileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}

	if len(req.Contract) == 0 {
		return errNoContractData
	} else {
		dataFiles, codeErr := codeReader(c)
		if codeErr != nil {
			//error reading source code
			return api.Error(c, codeErr)
		} else {
			//decoding success. compile the contract files
			compilerResponse, err := solc.CompileSolidity(dataFiles...)
			if err != nil {
				// compilation error
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

func (ctl SolcController) singleRawFileReader(c *echo.Context) ([]string, error) {
	//read request parameters encoded in the body
	req := protocol.SingleFileContractCompileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return noContractData, err
	}

	if req.Contract == "" {
		return noContractData, errNoContractData
	} else {
		return []string{req.Contract}, nil
	}
}

func (ctl SolcController) singleBase64FileReader(c *echo.Context) ([]string, error) {
	//read request parameters encoded in the body
	req := protocol.SingleFileContractCompileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return noContractData, err
	}

	if req.Contract == "" {
		return noContractData, errNoContractData
	} else {
		//decode contract from base64 string to ascii
		decoded, b64Err := base64.StdEncoding.DecodeString(req.Contract)
		if b64Err != nil {
			//b64 decoding error found
			return noContractData, b64Err
		} else {
			return []string{str.UnsafeString(decoded)}, nil
		}
	}
}

func (ctl SolcController) gitCodeReader(c *echo.Context) ([]string, error) {
	// Filesystem abstraction based on memory
	fs := memfs.New()
	// Git objects storer based on memory
	storer := memory.NewStorage()

	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
		/*Auth: &http.BasicAuth{
			Username: "username", // anything except an empty string
			Password: "password",
		},*/
	})
	if err != nil {
		return noContractData, err
	}

	var repoFiles []string
	dirFiles, err := fs.ReadDir(".")
	if err != nil {
		return noContractData, err
	} else {
		for _, f := range dirFiles {
			repoFiles = append(repoFiles, f.Name())
		}
	}
	return repoFiles, nil
}

func (ctl SolcController) uploadedZipReader(c *echo.Context) ([]string, error) {
	// source zipped file
	file, err := c.FormFile("file")
	if err != nil {
		return noContractData, err
	}
	src, err := file.Open()
	if err != nil {
		return noContractData, err
	}
	zippedBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return noContractData, err
	}
	_ = src.Close()
	files, err := packers.Unzip(zippedBytes)
	return files, err
}

func (ctl SolcController) uploadedTarGzReader(c *echo.Context) ([]string, error) {
	return []string{}, nil
}

func (ctl SolcController) compileModeSelector(c *echo.Context) error {
	mode := c.Param("mode")
	mode = str.ToLowerAscii(mode)
	logger.Debug("compiling solidity contract with mode: ", mode)
	switch mode {
	case data.SingleRawFile:
		logger.Debug("compiling ethereum contract from raw solidity source code")
		return ctl.compileFromSources(c, ctl.singleRawFileReader)
	case data.SingleBase64File:
		logger.Debug("compiling ethereum contract from solidity encoded single base64 source code")
		return ctl.compileFromSources(c, ctl.singleBase64FileReader)
	case data.GitMode:
		logger.Debug("compiling ethereum contract from remote git repository")
		return ctl.compileFromSources(c, ctl.gitCodeReader)
	case data.ZipMode:
		logger.Debug("compiling ethereum contract from user provided zip file")
		return ctl.compileFromSources(c, ctl.uploadedZipReader)
	case data.TargzMode:
		logger.Debug("compiling ethereum contract from user provided targz file")
		return ctl.compileFromSources(c, ctl.uploadedTarGzReader)
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
