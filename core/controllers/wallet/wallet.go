// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package wallet

import (
	"crypto/sha512"
	"encoding/hex"
	"strconv"

	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/go-hpc/lib/eth"
	"github.com/zerjioang/go-hpc/util/str"

	"github.com/zerjioang/go-hpc/lib/bip32"

	"github.com/zerjioang/etherniti/core/api"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/lib/bip39"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
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
func (ctl WalletController) Mnemonic(c *shared.EthernitiContext) error {

	req := dto.MnemonicRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error error
		logger.Error("failed to bind request data to model:", err)
		return api.ErrorBytes(c, data.BindErr)
	}

	// lowercase language
	req.Language = str.ToLowerAscii(req.Language)
	switch req.Language {
	case bip39.ChineseSimplified,
		bip39.ChineseTraditional,
		bip39.English,
		bip39.French,
		bip39.Italian,
		bip39.Japanese,
		bip39.Korean,
		bip39.Spanish:
		bip39.SetWordList(req.Language)
	default:
		//return invalid language error
		return api.ErrorBytes(c, data.MnemonicLanguageNotProvided)
	}

	if req.Size != 128 &&
		req.Size != 160 &&
		req.Size != 192 &&
		req.Size != 224 &&
		req.Size != 256 {
		//return invalid size error
		return api.ErrorBytes(c, data.MnemonicSizeNotSupported)
	}

	// create new Entropy from rand reader
	// Entropy is measured as bits and size measures bytes
	entropyBytes, entropyErr := bip39.GenerateSecureEntropy(req.Size)
	if entropyErr != nil {
		//failed to get a full Entropy source
		return api.ErrorBytes(c, data.InvalidEntropySource)
	}

	// create Mnemonic based on user config and created Entropy source
	mnemomic, err := bip39.NewMnemonic(entropyBytes)
	if err.Occur() {
		//return mnemonic error
		return api.StackError(c, err)
	} else {
		// hash the seed if requested
		var response dto.MnemonicResponse
		response.Language = req.Language
		response.Size = req.Size
		response.IsEncrypted = req.Secret != ""
		response.Mnemonic = mnemomic
		if response.IsEncrypted {
			// clear plaintext mnemonic
			response.Mnemonic = ""
			encryptedSeed := bip39.NewSeed(mnemomic, req.Secret)
			response.EncryptedSeed = hex.EncodeToString(encryptedSeed)
		}
		return api.SendSuccess(c, data.MnemonicSuccess, response)
	}
}

func (ctl WalletController) HdWallet(c *shared.EthernitiContext) error {
	req := dto.NewHdWalletRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model:", err)
		return api.ErrorBytes(c, data.BindErr)
	}
	response, err := ctl.createHdWallet(req)
	if err != nil {
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, data.HDWalletSuccess, response)
	}
}

func (ctl WalletController) Entropy(c *shared.EthernitiContext) error {
	req := dto.EntropyRequest{}
	req.Size = ctl.getIntParam(c, "bits")

	if req.Size <= 0 || req.Size > 4096*8 {
		//return invalid size (exceeded btw) error
		return api.ErrorBytes(c, data.EntropySizeNotSupported)
	}

	response, err := ctl.generateSecureEntropy(req)
	if err != nil {
		return api.Error(c, err)
	} else {
		//success
		return api.SendSuccess(c, data.EntropySuccess, response)
	}
}

func (ctl WalletController) generateSecureEntropy(request dto.EntropyRequest) (dto.EntropyResponse, error) {
	raw, err := bip39.GenerateSecureEntropy(request.Size)
	var response dto.EntropyResponse
	response.Raw = raw
	return response, err
}

func (ctl WalletController) createHdWallet(request dto.NewHdWalletRequest) (dto.HdWalletResponse, error) {
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

	var response dto.HdWalletResponse
	response.MasterPrivateKey = masterKey.String()
	response.MasterPublicKey = publicKey.String()
	response.Mnemonic = mnemonic
	return response, nil
}

// generates an ethereum new account (address+key)
func (ctl WalletController) newWallet(c *shared.EthernitiContext) error {

	// Create an account
	private, err := eth.GenerateNewKey()

	if err != nil {
		logger.Error("failed to generate ethereum account key", err)
		// send invalid generation message
		return api.ErrorBytes(c, data.EthAccountFailed)
	}
	address := eth.GetAddressFromPrivateKey(private)
	privateKey := eth.GetPrivateKeyAsEthString(private)
	var response dto.AccountResponse
	response.Address = address.Hex()
	response.Key = privateKey
	return api.SendSuccess(c, data.EthAccountSuccess, response)
}

// check if an ethereum address is valid
func (ctl WalletController) isValidAddress(c *shared.EthernitiContext) error {
	//since this method checks address as string, cache always
	c.OnSuccessCachePolicy = constants.CacheInfinite

	//read user entered address
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result := eth.IsValidAddressLow(targetAddr)
		return api.SendSuccess(c, data.EthAddressValidation, result)
	}
	// send invalid address message
	return api.ErrorBytes(c, data.MissingAddress)
}

func (ctl WalletController) getIntParam(c *shared.EthernitiContext, key string) uint16 {
	v := c.Param(key)
	if v != "" {
		num, _ := strconv.Atoi(v)
		return uint16(num)
	}
	return 0
}

// implemented method from interface RouterRegistrable
func (ctl WalletController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing wallet controller methods")
	router.POST("/wallet", wrap.Call(ctl.newWallet))
	router.POST("/hdwallet", wrap.Call(ctl.HdWallet))

	router.POST("/wallet/mnemonic/bip39", wrap.Call(ctl.Mnemonic))
	router.GET("/wallet/verify/:address", wrap.Call(ctl.isValidAddress))
	router.GET("/wallet/entropy/:bits", wrap.Call(ctl.Entropy))
}
