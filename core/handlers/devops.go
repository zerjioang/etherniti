package handlers

/*
// https://medium.com/@akshay_111meher/creating-offline-raw-transactions-with-go-ethereum-8d6cc8174c5d
// sends new eth transaction using given configuration
func (ctl Web3Controller) DeployContract(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, compiledBytecode string, contractAbiStr string) (string, error) {
	// Construct the transaction
	d := time.Now().Add(1000 * time.Millisecond)
	cancellableCtx, cancel := context.WithDeadline(ctx, d)

	// connect the client
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		logger.Error(err)
		cancel()
		return "", err
	}

	// remove 0x
	bytecode := common.Hex2Bytes(compiledBytecode[2:])
	contractAbi, err := abi.JSON(strings.NewReader(contractAbiStr))
	input, err := contractAbi.Pack("")

	bytecode = append(bytecode, input...)

	// To create a raw transaction nonce is must.
	// Hence we get the account nonce using
	// The third parameter is nil because we want the latest account nonce
	// To get nonce at a particular block height
	nonce, err := client.NonceAt(
		cancellableCtx,
		common.HexToAddress("0x56724a9e4d2bb2dca01999acade2e88a92b11a9e"), nil)


			The arguments to this function in order are
		    1. nonce
		    2. to-address (we would need to convert this using common.HexToAddress(public address) )
		    3. balance to be sent (use big number)
		    4. gas limit
		    5. gas price
		    6. data (since this is not a contract transaction, we can pass nil )

tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, bytecode)
// Define signer and chain id
// chainID := big.NewInt(CHAIN_ID)
// signer := types.NewEIP155Signer(chainID)
// FrontierSigner
// EIP155Signer
signer := types.HomesteadSigner{}

//create example account
privateKey, err := eth.GenerateNewKey()
// Sign the transaction signature with the private key
signedTx, signatureErr := types.SignTx(tx, signer, privateKey)
if signatureErr != nil {
fmt.Println("signature create trycatch:")
logger.Error(signatureErr)
cancel()
return "", signatureErr
}
// Send the signed transaction to the network
txErr := client.SendTransaction(cancellableCtx, signedTx)
if txErr != nil {
fmt.Println("send tx trycatch:")
logger.Error(txErr)
cancel()
return "", txErr
} else {
select {
case <-time.After(1 - time.Millisecond):
logger.Info("tx send overslept")
case <-cancellableCtx.Done():
logger.Info(cancellableCtx.Err())
default:
fmt.Printf("send success tx.hash=%s\n", signedTx.Hash().String())
cancel()
return signedTx.Hash().String(), nil
}
}
cancel()
return "", nil
}

// sends new eth transaction using given configuration
func (ctl Web3Controller) SendTransaction(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (string, error) {
	// Construct the transaction
	d := time.Now().Add(1000 * time.Millisecond)
	cancellableCtx, cancel := context.WithDeadline(ctx, d)

	// connect the client
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		logger.Error(err)
		cancel()
		return "", err
	}

	// To create a raw transaction nonce is must.
	// Hence we get the account nonce using
	// The third parameter is nil because we want the latest account nonce
	// To get nonce at a particular block height
	nonce, err := client.NonceAt(
		cancellableCtx,
		common.HexToAddress("0x56724a9e4d2bb2dca01999acade2e88a92b11a9e"), nil)


			The arguments to this function in order are
		    1. nonce
		    2. to-address (we would need to convert this using common.HexToAddress(public address) )
		    3. balance to be sent (use big number)
		    4. gas limit
		    5. gas price
		    6. data (since this is not a contract transaction, we can pass nil )

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	// Define signer and chain id
	// chainID := big.NewInt(CHAIN_ID)
	// signer := types.NewEIP155Signer(chainID)
	// FrontierSigner
	// EIP155Signer
	signer := types.HomesteadSigner{}

	//create example account
	privateKey, err := eth.GenerateNewKey()
	// Sign the transaction signature with the private key
	signedTx, signatureErr := types.SignTx(tx, signer, privateKey)
	if signatureErr != nil {
		fmt.Println("signature create trycatch:")
		logger.Error(signatureErr)
		cancel()
		return "", signatureErr
	}
	// Send the signed transaction to the network
	txErr := client.SendTransaction(cancellableCtx, signedTx)
	if txErr != nil {
		fmt.Println("send tx trycatch:")
		logger.Error(txErr)
		cancel()
		return "", txErr
	} else {
		select {
		case <-time.After(1 - time.Millisecond):
			logger.Info("tx send overslept")
		case <-cancellableCtx.Done():
			logger.Info(cancellableCtx.Err())
		default:
			fmt.Printf("send success tx.hash=%s\n", signedTx.Hash().String())
			cancel()
			return signedTx.Hash().String(), nil
		}
	}
	cancel()
	return "", nil
}

// call deployed smart contract method
func (ctl Web3Controller) CallContract() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		logger.Error(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		logger.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("trycatch casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		logger.Error(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		logger.Error(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

}
*/
