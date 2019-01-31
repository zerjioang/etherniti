// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"github.com/zerjioang/gaethway/core/modules/ethfork/common"

	"github.com/zerjioang/gaethway/core/modules/ethfork/common/hexutil"
	"github.com/zerjioang/gaethway/core/modules/ethfork/crypto"
	"github.com/zerjioang/gaethway/core/modules/ethfork/ethclient"
)

type EthClientType uint8

const (
	HttpClient EthClientType = iota
	IPCClient
)

type EthereumClient *ethclient.Client

// get an ethereum client
// ethGateway is the geth client running or
// infura like system: https://mainnet.infura.io
// ganache like: http://localhost:8545
func getClient(ethGateway string) (*ethclient.Client, error) {
	return ethclient.Dial(ethGateway)
}

// get an ethereum client using ipc communication
func getIPCClient(ipcEndpoint string) (*ethclient.Client, error) {
	return ethclient.Dial(ipcEndpoint)
}

// get an ethereum client using specified mode and gateway
func GetEthereumClient(mode EthClientType, gateway string) (*ethclient.Client, error) {
	if mode == HttpClient {
		return getClient(gateway)
	}
	if mode == IPCClient {
		return getIPCClient(gateway)
	}
	return nil, errors.New("invalid mode")
}

// generates new ethereum ECDSA key
// This is the private key which is used for signing transactions and is to be treated
// like a password and never be shared, since who ever is in possesion
// of it will have access to all your funds.
func GenerateNewKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

// get the bytes of private key
func GetPrivateKeyBytes(privateKey *ecdsa.PrivateKey) []byte {
	return crypto.FromECDSA(privateKey)
}

// get the private key encoded as 0x..... string
func GetPrivateKeyAsEthString(privateKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(privateKey.D.Bytes())
}

// get the bytes of private key
func GetPublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return privateKey.Public().(*ecdsa.PublicKey)
}

// get ethereum address from its private key
func GetAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

// get the bytes of public key
func GetPublicKeyBytes(pub *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(pub)
}

// The public address is simply the Keccak-256 hash of the public key
// and then we take the last 40 characters (20 bytes) and prefix it with 0x
func GetPublicKeyAsEthString(pub *ecdsa.PublicKey) string {
	return hexutil.Encode(
		GetPublicKeyBytes(pub),
	)[4:]
}
