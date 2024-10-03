package main

import (
        "context"
        "fmt"
        "log"
        "math/big"

        "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
        client, err := ethclient.Dial("https://cloudflare-eth.com")
        if err != nil {
                log.Fatal(err)
        }

        //调用返回最新区块的头部信息
        //区块头包含了区块的元数据，如区块号、时间戳、难度等，但不包括区块体（即交易列表）
        header, err := client.HeaderByNumber(context.Background(), nil)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(header.Number.String()) // 5671744

        //获取特定区块的完整信息
        blockNumber := big.NewInt(5671744)
        block, err := client.BlockByNumber(context.Background(), blockNumber)
        if err != nil {
                log.Fatal(err)
        }

        //输出区块的详细信息 输出区块的编号、时间戳、难度、哈希值和交易数量
        fmt.Println(block.Number().Uint64())     // 5671744
        fmt.Println(block.Time())       // 1527211625
        fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
        fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
        fmt.Println(len(block.Transactions()))   // 144

        //获取特定区块的交易数量
        count, err := client.TransactionCount(context.Background(), block.Hash())
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(count) // 144
}