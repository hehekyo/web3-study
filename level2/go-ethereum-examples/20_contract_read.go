package main

import (
        "fmt"
        "log"

        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/ethclient"

        store "go-ethereum-examples/contracts/store"
)

func main() {
        client, err := ethclient.Dial("https://sepolia.infura.io/v3/a1439f90e4fe49b0a48be70256fe8af1")
        if err != nil {
                log.Fatal(err)
        }

        address := common.HexToAddress("0x61774B7094Bc76bDf90c67b4f1aE66B00d1556B4")
        instance, err := store.NewStore(address, client)
        if err != nil {
                log.Fatal(err)
        }

        version, err := instance.Version(nil)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(version) // "1.0"
}