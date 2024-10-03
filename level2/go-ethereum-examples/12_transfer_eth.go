package main

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // 连接到 Ganache
    // client, err := ethclient.Dial("http://localhost:7545")
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/a1439f90e4fe49b0a48be70256fe8af1")

    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // Ganache 默认账户私钥（从 Ganache UI 获取）
    privateKey, err := crypto.HexToECDSA("1c7bd28c7b452b76277b82e3aaddb90fbdbacf80e4ba45a5b9548c4dbb18ff9e")
    if err != nil {
        log.Fatal("HexToECDSA:",err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    fmt.Printf("From Address: %s\n", fromAddress.Hex())
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    // value := big.NewInt(1000000000000000000) // 1 ETH
    value := big.NewInt(1000000000000000) // 0.001 ETH
    gasLimit := uint64(21000)               // ETH 转账的标准 gas limit
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    // 接收者地址
    toAddress := common.HexToAddress("0x7EA4F504727E86d77D7eEB679e91c760D6ba667D")
    var data []byte // 转账交易的 data 字段为空
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
    //print nonce, toAddress, value, gasLimit, gasPrice, data
    fmt.Printf(" Nonce: %d\n To Address: %s\n Value: %s\n Gas Limit: %d\n Gas Price: %s\n Data: %x\n",nonce, toAddress, value, gasLimit, gasPrice, data)
    fmt.Printf("Transaction details: %+v\n", tx)

   // chainID, err := client.NetworkID(context.Background())   this is wrong
    chainID, err := client.ChainID(context.Background())
    fmt.Printf("Chain ID: %d\n", chainID)
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal("signedTx:",err)
    }

    fmt.Printf("Signed transaction details: %+v\n", signedTx)

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal("SendTransaction:",err)
    }

    //转账成功,在这里可以查到
    //https://sepolia.etherscan.io/tx/0x0004a2598d4bc0a611a32a0b4e3630898156d00760af4bf12df050b5d58b965c
    fmt.Printf("Transaction sent! TX Hash: %s\n", signedTx.Hash().Hex())
}
