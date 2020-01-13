package network

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/pkg/errors"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/network/signer"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/lib/eth/fixtures/common"
	"github.com/zerjioang/go-hpc/lib/eth/fixtures/crypto"
	ethrpc "github.com/zerjioang/go-hpc/lib/eth/rpc"
)

func (ctl *Web3Controller) signTransactionLocal(c *shared.EthernitiContext) error {
	// 0 parse this http request body
	var req *dto.EthSignRequest
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to req: ", err)
		return api.ErrorBytes(c, data.BindErr)
	}
	err := req.Validate()
	if err == nil {

		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		// execute tx signing process
		signedTx, txSignErr := localSignTransaction(client, req)
		if txSignErr != nil {
			return api.Error(c, txSignErr)
		} else {
			return api.SendSuccess(c, data.EthLocalSign, signedTx)
		}
	} else {
		// error validating input data
		logger.Error("error validating input data: ", err)
		return api.Error(c, err)
	}
}

func localSignTransaction(client *ethrpc.EthRPC, req *dto.EthSignRequest) (string, error) {
	// get our client context
	if client == nil {
		return "", errors.New("missing client for transaction signing")
	}
	if req == nil {
		return "", errors.New("missing transaction signing request to be signed")
	}
	valErr := req.Validate()
	if valErr != nil {
		return "", valErr
	}

	privateKey, err := crypto.HexToECDSA(req.PrivateKey)
	if err != nil {
		logger.Error("failed to decode private key from hexadecimal encoded string")
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("error casting public key to ECDSA")
		return "", errors.New("invalid ECDSA key provided")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fromStr := fromAddress.Hex()
	if req.From != fromStr {
		inconsistentErr := errors.New("from field is not valid. consistency with private key broken")
		logger.Error(inconsistentErr)
		return "", inconsistentErr
	}
	nonce, err := client.PendingNonceAt(req.From)
	if err != nil {
		logger.Error("failed to get pending nonce ethValue at current time due to: ", err)
		return "", err
	}

	ethValue := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                   // in units
	gasPrice, err := client.SuggestGasPrice()
	if err != nil {
		logger.Error("failed to get current blockchain suggested gas price due to: ", err)
		return "", err
	}

	toAddress := common.HexToAddress(req.To)
	var payload []byte
	tx := signer.NewTransaction(nonce, toAddress, ethValue, gasLimit, gasPrice, payload)

	chainID, err := client.NetworkId()
	if err != nil {
		logger.Error("failed to get blockchain network id due to: ", err)
		return "", err
	}

	// Define signer and chain id
	// chainID := big.NewInt(CHAIN_ID)
	// signer := types.NewEIP155Signer(chainID)
	signerModule := signer.NewEIP155Signer(chainID)

	signedTx, err := signer.SignTx(tx, signerModule, privateKey)
	if err != nil {
		logger.Error("failed to get sign given transaction: ", err)
		return "", err
	}

	ts := signer.Transactions{signedTx}
	rawTx := hex.EncodeToString(ts.GetRlp(0))
	return rawTx, nil
}
