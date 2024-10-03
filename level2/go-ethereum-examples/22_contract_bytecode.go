package main

import (
        "context"
        "encoding/hex"
        "fmt"
        "log"

        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
        client, err := ethclient.Dial("https://sepolia.infura.io/v3/a1439f90e4fe49b0a48be70256fe8af1")
        if err != nil {
                log.Fatal(err)
        }

        contractAddress := common.HexToAddress("0x61774B7094Bc76bDf90c67b4f1aE66B00d1556B4")
        bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}