// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package solc

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/lib/packers"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"gopkg.in/src-d/go-billy.v4/memfs"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/lib/solc"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
	"github.com/zerjioang/go-hpc/util/str"
)

// token controller
type SolcController struct {
}

type CodeReader func(c echo.Context) (*dto.ContractCompilationOpts, []string, error)

var (
	solcResponseCode          codes.HttpStatusCode
	solcVersionErr            error
	solcVersionData           *solc.Solidity
	errNoContractData         = errors.New("failed to get request contract data")
	errNoSolc                 = errors.New("failed to get solc version")
	noContractData            []string
	defaultCompilationOptions = &dto.ContractCompilationOpts{
		EvmVersion: "petersburg",
		Optimize: dto.OptimizeOpts{
			Enabled: true,
			Runs:    200,
		},
		Gas:     true,
		Machine: "evm",
		Report: dto.ReportOpts{
			Opcodes:    true,
			Bin:        true,
			BinRuntime: true,
			Hashes:     true,
		},
	}
)

func init() {
	//preload solc version information
	solData, err := solc.SolidityVersion()
	if solData == nil {
		solcResponseCode = codes.StatusBadRequest
		solcVersionErr = errNoSolc
	} else if err != nil {
		solcResponseCode = codes.StatusBadRequest
		solcVersionErr = err
	} else {
		solcResponseCode = codes.StatusOK
		solcVersionData = solData
	}
}

// constructor like function
func NewSolcController() SolcController {
	return SolcController{}
}

func (ctl SolcController) version(c *shared.EthernitiContext) error {
	if solcResponseCode == 200 {
		return api.SendSuccess(c, data.SolcVersion, solcVersionData)
	} else {
		return api.ErrorCode(c, solcResponseCode, solcVersionErr)
	}
}

// solc --optimize --optimize-runs 200 --opcodes --bin --abi --hashes --asm erc20.sol
func (ctl SolcController) compileFromSources(c *shared.EthernitiContext, codeReader CodeReader) error {
	// 1 read user solc configuration
	// 2 read user source code
	opts, dataFiles, codeErr := codeReader(c)
	if opts == nil {
		//use default compilation options
		opts = defaultCompilationOptions
	}
	if codeErr != nil {
		//error reading source code
		return api.Error(c, codeErr)
	} else {
		//decoding success. compile the contract files
		compilerResponse, err := solc.CompileSolidityFileBytes(dataFiles)
		if err != nil {
			// compilation error
			return api.Error(c, err)
		} else {
			//generate the response for the client
			var contractResp []dto.ContractCompileResponse
			contractResp = make([]dto.ContractCompileResponse, len(compilerResponse))
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

func (ctl SolcController) singleRawFileReader(c echo.Context) (*dto.ContractCompilationOpts, []string, error) {
	//read request parameters encoded in the body
	req := dto.SingleFileContractCompileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return &req.Opts, noContractData, err
	}

	if req.Contract == "" {
		return &req.Opts, noContractData, errNoContractData
	} else {
		return &req.Opts, []string{req.Contract}, nil
	}
}

func (ctl SolcController) singleBase64FileReader(c echo.Context) (*dto.ContractCompilationOpts, []string, error) {
	//read request parameters encoded in the body
	req := dto.SingleFileContractCompileRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return &req.Opts, noContractData, err
	}

	if req.Contract == "" {
		return &req.Opts, noContractData, errNoContractData
	} else {
		//decode contract from base64 string to ascii
		decoded, b64Err := base64.StdEncoding.DecodeString(req.Contract)
		if b64Err != nil {
			//b64 decoding error found
			return &req.Opts, noContractData, b64Err
		} else {
			return &req.Opts, []string{str.UnsafeString(decoded)}, nil
		}
	}
}

func (ctl SolcController) gitCodeReader(c echo.Context) (*dto.ContractCompilationOpts, []string, error) {

	// 1 read compilation settings
	var opts *dto.ContractCompilationOpts
	if err := c.Bind(&opts); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return nil, noContractData, err
	}

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
		return opts, noContractData, err
	}

	var repoFiles []string
	dirFiles, err := fs.ReadDir(".")
	if err != nil {
		return opts, noContractData, err
	} else {
		for _, f := range dirFiles {
			repoFiles = append(repoFiles, f.Name())
		}
	}
	return opts, repoFiles, nil
}

func (ctl SolcController) uploadedZipReader(c echo.Context) (*dto.ContractCompilationOpts, []string, error) {

	// 1 read compilation settings
	var opts *dto.ContractCompilationOpts
	if err := c.Bind(&opts); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return nil, noContractData, err
	}

	// source zipped file
	file, err := c.FormFile("file")
	if err != nil {
		return opts, noContractData, err
	}
	src, err := file.Open()
	if err != nil {
		return opts, noContractData, err
	}
	zippedBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return opts, noContractData, err
	}
	_ = src.Close()
	files, err := packers.Unzip(zippedBytes)
	return opts, files, err
}

func (ctl SolcController) uploadedTarGzReader(c echo.Context) (*dto.ContractCompilationOpts, []string, error) {

	// 1 read compilation settings
	var opts *dto.ContractCompilationOpts
	if err := c.Bind(&opts); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return nil, noContractData, err
	}

	return opts, []string{}, nil
}

func (ctl SolcController) compileModeSelector(c *shared.EthernitiContext) error {
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
	router.GET("/solc/version", wrap.Call(ctl.version))
	router.POST("/solc/compile/:mode", wrap.Call(ctl.compileModeSelector))
}
