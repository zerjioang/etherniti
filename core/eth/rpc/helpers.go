// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ethrpc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/core/eth/fixtures/crypto"
	"github.com/zerjioang/etherniti/core/logger"
	"golang.org/x/crypto/sha3"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
)

// ParseInt parse hex string value to int
func ParseInt(value string) (int, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

// ParseBigInt parse hex string value to big.Int
func ParseBigInt(value string) (*big.Int, error) {
	i := new(big.Int)
	i, _ = i.SetString(value, 16)
	return i, nil
}

// ParseBigInt parse hex string value to big.Int
func ParseHexToInt(value string) (int64, error) {
	return strconv.ParseInt(value, 16, 64)
}

// Int64ToHex convert int64 to hexadecimal representation
func Int64ToHex(i int64) string {
	return "0x" + strconv.FormatInt(i, 16)
}

// IntToHex convert int to hexadecimal representation
func IntToHex(i int) string {
	return Int64ToHex(int64(i))
}

// BigToHex covert big.Int to hexadecimal representation
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}

	return fixtures.Encode(bigInt.Bytes())
}

func HexToBigInt(hex string) *big.Int {
	i := new(big.Int)
	i, _ = i.SetString(hex, 16)
	return i
}

// TextAndHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func TextAndHash(data []byte) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(msg))
	return hasher.Sum(nil), msg
}

func LocalSigning(addrHex string, privHex string, plainMessage string) ([]byte, error) {
	key, cErr := crypto.HexToECDSA(privHex)
	if cErr != nil {
		logger.Error("private key load error: ", cErr)
		return nil, cErr
	}

	hashed, message := TextAndHash([]byte("foo-bar"))
	hashHexMsg := fixtures.Encode(hashed)
	logger.Info(string(hashed))
	logger.Info(hashHexMsg)
	logger.Info(message)

	sig, err := crypto.Sign(hashed, key)
	if err != nil {
		logger.Error("signing error: ", err)
		return nil, err
	}
	return sig, err
	//addr := fixtures.HexToAddress(addrHex)
	/*recoveredPub, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		logger.Error("eCRecover error: %s", err)
		return err
	}
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := fixtures.PubkeyToAddress(*pubKey)
	if addr != recoveredAddr {
		logger.Error("address mismatch: want: %x have: %x", addr, recoveredAddr)
		return err
	}

	// should be equal to SigToPub
	recoveredPub2, err := crypto.SigToPub(msg, sig)
	if err != nil {
		logger.Error("eCRecover error: %s", err)
		return err
	}
	recoveredAddr2 := fixtures.PubkeyToAddress(*recoveredPub2)
	if addr != recoveredAddr2 {
		logger.Error("address mismatch: want: %x have: %x", addr, recoveredAddr2)
		return err
	}*/
}
