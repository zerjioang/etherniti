// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ethrpc

/*
bzz_hive (stub)
bzz_info (stub)
debug_traceTransaction
eth_accounts
eth_blockNumber
eth_call
eth_coinbase
eth_estimateGas
eth_gasPrice
eth_getBalance
eth_getBlockByNumber
eth_getBlockByHash
eth_getBlockTransactionCountByHash
eth_getBlockTransactionCountByNumber
eth_getCode (only supports block number “latest”)
eth_getCompilers
eth_getFilterChanges
eth_getFilterLogs
eth_getLogs
eth_getStorageAt
eth_getTransactionByHash
eth_getTransactionByBlockHashAndIndex
eth_getTransactionByBlockNumberAndIndex
eth_getTransactionCount
eth_getTransactionReceipt
eth_hashrate
eth_mining
eth_newBlockFilter
eth_newFilter (includes log/event filters)
eth_protocolVersion
eth_sendTransaction
eth_sendRawTransaction
eth_sign
eth_subscribe (only for websocket connections. "syncing" subscriptions are not yet supported)
eth_unsubscribe (only for websocket connections. "syncing" subscriptions are not yet supported)
eth_syncing
eth_uninstallFilter
net_listening
net_peerCount
net_version
miner_start
miner_stop
personal_listAccounts
personal_lockAccount
personal_newAccount
personal_importRawKey
personal_unlockAccount
personal_sendTransaction
shh_version
rpc_modules
web3_clientVersion
web3_sha3

https://ethereumbuilders.gitbooks.io/guide/content/en/ethereum_json_rpc.html

package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    store "./contracts" // for demo
)

func main() {
    client, err := ethclient.Dial("https://rinkeby.infura.io")
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
    instance, err := store.NewStore(address, client)
    if err != nil {
        log.Fatal(err)
    }

    key := [32]byte{}
    value := [32]byte{}
    copy(key[:], []byte("foo"))
    copy(value[:], []byte("bar"))

    tx, err := instance.SetItem(auth, key, value)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

    result, err := instance.Items(nil, key)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(result[:])) // "bar"
}
*/
