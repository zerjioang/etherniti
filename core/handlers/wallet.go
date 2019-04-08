// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/core/handlers/clientcache"
	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/modules/bip32"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/bip39"
)

const (
	defaultPath      = "m/44'/60'/0'/0/0"
	invalidAddress   = `{"message": "please, provide a valid ethereum or quorum address"}`
	accountKeyGenErr = `{"message": "failed to generate ecdsa private key"}`
)

var (
	noConnErrMsg           = "invalid connection profile key provided in the request header. Please, make sure you have created a connection profile indicating your peer node IP address or domain name."
	errNoConnectionProfile = errors.New(noConnErrMsg)
	accountKeyGenErrBytes  = str.UnsafeBytes(accountKeyGenErr)
	invalidAddressBytes    = str.UnsafeBytes(invalidAddress)
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
func (ctl WalletController) Mnemonic(c echo.Context) error {

	req := protocol.MnemonicRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error error
		logger.Error("failed to bind request data to model:", err)
		return api.ErrorStr(c, bindErr)
	}

	// lowercase language
	req.Language = str.ToLowerAscii(req.Language)

	if req.Language == "chinese-simplified" {
		bip39.SetWordList(req.Language)
	} else if req.Language == "chinese-traditional" {
		bip39.SetWordList(req.Language)
	} else if req.Language == "english" {
		bip39.SetWordList(req.Language)
	} else if req.Language == "french" {
		bip39.SetWordList(req.Language)
	} else if req.Language == "italian" {
		bip39.SetWordList(req.Language)
	} else if req.Language == "japanese" {
		bip39.SetWordList(req.Language)
	} else if req.Language == "korean" {
		bip39.SetWordList(req.Language)
	} else if req.Language == "spanish" {
		bip39.SetWordList(req.Language)
	} else {
		//return invalid language error
		return api.ErrorStr(c, "provided language is not supported")
	}

	if req.Size != 128 &&
		req.Size != 160 &&
		req.Size != 192 &&
		req.Size != 224 &&
		req.Size != 256 {
		//return invalid size error
		return api.ErrorStr(c, "provided mnemonic size is not supported")
	}

	// create new Entropy from rand reader
	// Entropy is measured as bits and size measures bytes
	var entropyBytes = uint8(req.Size / 8)
	entropy := make([]byte, entropyBytes)
	n, readErr := io.ReadFull(rand.Reader, entropy)
	if readErr != nil || uint8(n) != entropyBytes {
		//failed to get a full Entropy source
		return api.ErrorStr(c, "failed to get a full Entropy source")
	}

	// create Mnemonic based on user config and created Entropy source
	mnemomic, err := bip39.NewMnemonic(entropy)
	if err.Occur() {
		//return mnemonic error
		return api.StackError(c, err)
	} else {
		// hash the seed if requested
		var response protocol.MnemonicResponse
		response.Language = req.Language
		response.Size = req.Size
		if req.Secret != "" {
			encryptedSeed := bip39.NewSeed(mnemomic, req.Secret)
			response.EncryptedSeed = string(encryptedSeed)
		}
		return api.SendSuccess(c, "mnemonic successfully created", response)
	}
}

func (ctl WalletController) HdWallet(c echo.Context) error {
	req := protocol.NewHdWalletRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model:", err)
		return api.ErrorStr(c, bindErr)
	}
	response, err := ctl.createHdWallet(req)
	if err != nil {
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, "hd wallet successfully created", response)
	}
}

func (ctl WalletController) Entropy(c echo.Context) error {
	req := protocol.EntropyRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model:", err)
		return api.ErrorStr(c, bindErr)
	}

	req.Size = ctl.getIntParam(c, "bits")

	if req.Size <= 0 || req.Size > 4096*8 {
		//return invalid size (exceeded btw) error
		return api.ErrorStr(c, "provided entropy size is not supported")
	}

	response, err := ctl.generateSecureEntropy(req)
	if err != nil {
		return api.Error(c, err)
	} else {
		//success
		return api.SendSuccess(c, "Entropy data generated", response)
	}
}

func (ctl WalletController) generateSecureEntropy(request protocol.EntropyRequest) (protocol.EntropyResponse, error) {
	raw, err := bip39.GenerateSecureEntropy(request.Size)
	var response protocol.EntropyResponse
	response.Raw = raw
	return response, err
}

func (ctl WalletController) createHdWallet(request protocol.NewHdWalletRequest) (protocol.HdWalletResponse, error) {
	// Generate a Mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the Mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed, "Bitcoin seed", sha512.New)
	publicKey := masterKey.PublicKey()

	// Display Mnemonic and keys
	logger.Debug("Mnemonic: ", mnemonic)
	logger.Debug("Master private key: ", masterKey)
	logger.Debug("Master public key: ", publicKey)

	var response protocol.HdWalletResponse
	response.MasterPrivateKey = masterKey.String()
	response.MasterPublicKey = publicKey.String()
	response.Mnemonic = mnemonic
	return response, nil
}

// generates an ethereum new account (address+key)
func (ctl WalletController) generateAddress(c echo.Context) error {

	// Create an account
	private, err := eth.GenerateNewKey()

	if err != nil {
		logger.Error("failed to generate ethereum account key", err)
		// send invalid generation message
		return c.JSONBlob(http.StatusInternalServerError, accountKeyGenErrBytes)
	}
	address := eth.GetAddressFromPrivateKey(private)
	privateKey := eth.GetPrivateKeyAsEthString(private)
	var response = map[string]string{
		"address": address.Hex(),
		"private": privateKey,
	}
	return c.JSONBlob(
		http.StatusOK,
		str.GetJsonBytes(
			protocol.NewApiResponse("ethereum account created", response),
		),
	)
}

// check if an ethereum address is valid
func (ctl WalletController) isValidAddress(c echo.Context) error {
	//since this method checks address as string, cache always
	var code int
	code, c = clientcache.Cached(c, true, clientcache.CacheInfinite) // 24h cache directive

	//read user entered address
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result := eth.IsValidAddress(targetAddr)
		return c.JSONBlob(code, str.GetJsonBytes(
			protocol.NewApiResponse("address validation checked", result),
		),
		)
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// implemented method from interface RouterRegistrable
func (ctl WalletController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing wallet controller methods")
	router.POST("/wallet", ctl.generateAddress)
	router.GET("/wallet/verify/:address", ctl.isValidAddress)
	router.GET("/wallet/entropy/:bits", ctl.Entropy)
	router.POST("/wallet/mnemonic/bip39", ctl.Mnemonic)
	router.POST("/wallet/hd/bip32", ctl.HdWallet)
}

func (ctl WalletController) getIntParam(c echo.Context, key string) uint16 {
	v := c.Param(key)
	if v != "" {
		num, _ := strconv.Atoi(v)
		return uint16(num)
	}
	return 0
}
