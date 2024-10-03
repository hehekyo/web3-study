package main

import (
        "context"
        "fmt"
        "log"
        "regexp"

        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
        //正则表达式匹配
        re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

        fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
        fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

        client, err := ethclient.Dial("https://cloudflare-eth.com")
        if err != nil {
                log.Fatal(err)
        }

        //检查智能合约代码

        // 0x Protocol Token (ZRX) smart contract address
        address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
        //client.CodeAt()检查逻辑
        //当调用 CodeAt 方法时，它会向连接的以太坊节点发起一个 JSON RPC 请求，具体是 eth_getCode 请求。
        //这个请求询问节点在指定区块高度的给定地址存储的智能合约代码。
        //如果地址是普通的用户钱包地址，它不会有任何关联的智能合约代码，因此返回的结果将是空的字节串。
        //如果地址是智能合约地址，则会返回该合约的字节码。
        bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
        if err != nil {
                log.Fatal(err)
        }
        println("bytecode:",bytecode) //bytecode: [2946/2946]0x1400041c000

        isContract := len(bytecode) > 0

        fmt.Printf("is contract: %v\n", isContract) // is contract: true

        // a random user account address
        address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
        bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
        if err != nil {
                log.Fatal(err)
        }

        isContract = len(bytecode) > 0

        fmt.Printf("is contract: %v\n", isContract) // is contract: false
}