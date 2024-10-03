package main

import (
        "context"
        "fmt"
        "log"

        "github.com/ethereum/go-ethereum"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/types"
        "github.com/ethereum/go-ethereum/ethclient"
)

//订阅事件日志

func main() {
        //拨打启用 websocket 的以太坊客户端
        client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/a1439f90e4fe49b0a48be70256fe8af1")
        if err != nil {
                log.Fatal(err)
        }

        contractAddress := common.HexToAddress("0x61774B7094Bc76bDf90c67b4f1aE66B00d1556B4")
        query := ethereum.FilterQuery{
                Addresses: []common.Address{contractAddress},
        }

        logs := make(chan types.Log)
        sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
        if err != nil {
                log.Fatal(err)
        }

        for {
                select {
                case err := <-sub.Err():
                        log.Fatal(err)
                case vLog := <-logs:
                        fmt.Println(vLog) // pointer to event log
                }
        }
}