package main

import (
        "context"
        "encoding/hex"
        "fmt"
        "log"

        "github.com/ethereum/go-ethereum/core/types"
        "github.com/ethereum/go-ethereum/ethclient"
        "github.com/ethereum/go-ethereum/rlp"
)

/**
发送原始交易事务
**/
func main() {
        client, err := ethclient.Dial("https://sepolia.infura.io/v3/a1439f90e4fe49b0a48be70256fe8af1")
        if err != nil {
                log.Fatal(err)
        }

        rawTx := "f86f05852993f5ae50825208947ea4f504727e86d77d7eeb679e91c760d6ba667d872386f26fc10000808401546d72a0a614a80fc0b0555d82320dc7a56e7fe398e4c29ab541049a504763af07314392a04374c85c29791df072fe4c8d6eea818320b2b6de471da79104b3588f6dcf3916"

        rawTxBytes, err := hex.DecodeString(rawTx)

        tx := new(types.Transaction)
        rlp.DecodeBytes(rawTxBytes, &tx)

        err = client.SendTransaction(context.Background(), tx)
        if err != nil {
                log.Fatal(err)
        }

        //https://sepolia.etherscan.io/tx/0xf23a00119f40c105791acca88f986eb982019a6bc79e61e4254e0b6414e5d880
        fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xf23a00119f40c105791acca88f986eb982019a6bc79e61e4254e0b6414e5d880
}