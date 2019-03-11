// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"crypto/rand"
	"io"
	"net/http"

	"github.com/zerjioang/etherniti/core/modules/bip32"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/bip39"
	"github.com/zerjioang/etherniti/core/modules/bip39/wordlists"
	"github.com/zerjioang/etherniti/core/util"
)

const (
	defaultPath = "m/44'/60'/0'/0/0"
)

type WalletController struct {
}

func NewWalletController() WalletController {
	dc := WalletController{}
	return dc
}

// generate a mnemomic
/*
	128 bits -> 12 words
	160 bits -> 15 words
	192 bits -> 18 words
	224 bits -> 21 words
	256 bits -> 24 words
*/
func (ctl WalletController) mnemonic(c echo.Context) error {

	req := protocol.NewMnemonicRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding trycatch error
		logger.Error("failed to bind request data to model:", err)
		return api.ErrorStr(c, bindErr)
	}

	// lowercase language
	req.Language = util.ToLowerAscii(req.Language)

	if req.Language == "chinese-simplified" {
		bip39.SetWordList(wordlists.ChineseSimplified)
	} else if req.Language == "chinese-traditional" {
		bip39.SetWordList(wordlists.ChineseTraditional)
	} else if req.Language == "english" {
		bip39.SetWordList(wordlists.English)
	} else if req.Language == "french" {
		bip39.SetWordList(wordlists.French)
	} else if req.Language == "italian" {
		bip39.SetWordList(wordlists.Italian)
	} else if req.Language == "japanese" {
		bip39.SetWordList(wordlists.Japanese)
	} else if req.Language == "korean" {
		bip39.SetWordList(wordlists.Korean)
	} else if req.Language == "spanish" {
		bip39.SetWordList(wordlists.Spanish)
	} else {
		//return invalid language trycatch
		return api.ErrorStr(c, "provided language is not supported")
	}

	if req.Size != 128 &&
		req.Size != 160 &&
		req.Size != 192 &&
		req.Size != 224 &&
		req.Size != 256 {
		//return invalid size trycatch
		return api.ErrorStr(c, "provided size is not supported")
	}

	// create new entropy from rand reader
	// entropy is measured as bits and size measures bytes
	var entropyBytes = uint8(req.Size / 8)
	entropy := make([]byte, entropyBytes)
	n, readErr := io.ReadFull(rand.Reader, entropy)
	if readErr != nil || uint8(n) != entropyBytes {
		//failed to get a full entropy source
		return api.ErrorStr(c, "failed to get a full entropy source")
	}

	// create mnemonic based on user config and created entropy source
	mnemomic, err := bip39.NewMnemonic(entropy)
	if err.Occur() {
		//return mnemonic error
		return api.StackError(c, err)
	} else {
		//return mnemonic content
		rawBytes := util.GetJsonBytes(protocol.NewApiResponse("mnemonic successfully created", mnemomic))
		return c.JSONBlob(http.StatusOK, rawBytes)
	}
}

func (ctl WalletController) hdWallet(c echo.Context) error {
	req := protocol.NewHdWalletRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding trycatch error
		logger.Error("failed to bind request data to model:", err)
		return api.ErrorStr(c, bindErr)
	}
	response, err := ctl.createHdWallet(req)
	if err != nil {
		return api.Error(c, err)
	} else {
		//hd wallet success
		logger.Info("hd wallet successfully created")
		response := api.ToSuccess("hd wallet successfully created", response)
		return c.JSONBlob(protocol.StatusOK, response)
	}
}

func (ctl WalletController) createHdWallet(request protocol.NewHdWalletRequest) (protocol.HdWalletResponse, error) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	logger.Debug("Mnemonic: ", mnemonic)
	logger.Debug("Master private key: ", masterKey)
	logger.Debug("Master public key: ", publicKey)

	var response protocol.HdWalletResponse
	response.MasterPrivateKey = masterKey.String()
	response.MasterPublicKey = publicKey.String()
	response.Mnemonic = mnemonic
	return response, nil
}

// implemented method from interface RouterRegistrable
func (ctl WalletController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing wallet controller methods")
	router.POST("/mnemonic/bip39", ctl.mnemonic)
	router.POST("/hd/bip32", ctl.hdWallet)
}
