package main

import (
        "context"
        "fmt"
        "log"
        "math/big"

        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/types"
        "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
        client, err := ethclient.Dial("https://cloudflare-eth.com")
        if err != nil {
                log.Fatal(err)
        }

        blockNumber := big.NewInt(5671744)
        block, err := client.BlockByNumber(context.Background(), blockNumber)
        if err != nil {
                log.Fatal(err)
        }

        // 遍历区块中的所有交易 只获取 1 个交易
        // 可以在这里查看 https://etherscan.io/tx/0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
        for _, tx := range block.Transactions()[:1] {
                fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
                fmt.Println(tx.Value().String())    // 10000000000000000
                fmt.Println(tx.Gas())               // 105000 消耗的 Gas
                fmt.Println(tx.GasPrice().Uint64()) // 102000000000 Gas 价格
                fmt.Println(tx.Nonce())             // 110644
                fmt.Println(tx.Data())              // [] 交易数据
                fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e 接收地址

                //当前网络的 chain ID
                chainID, err := client.NetworkID(context.Background())
                if err != nil {
                        log.Fatal(err)
                }

                //解析出交易的发送者地址
                if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
                        fmt.Println("sender", sender.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
}

                //获取交易的收据
                receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
                if err != nil {
                        log.Fatal(err)
                }
                // 交易收据包含了交易执行的结果状态，如成功或失败，以及该交易影响的日志事件等信息。
                // 交易状态(Status) Gas Used Cumulative Gas Used 交易日志
                // 查看交易日志 https://etherscan.io/tx/0x50a55e165a5db0074e32556bb0473aff747a5c0794c007a3b333c9d4eef7b70c#eventlog
                fmt.Println(receipt.Status) // 1
        }

        //将这个十六进制字符串转换成以太坊 common.Hash 类型
        blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
        count, err := client.TransactionCount(context.Background(), blockHash)
        if err != nil {
                log.Fatal(err)
        }

        for idx := uint(0); idx < count; idx++ {
                //获取区块中指定索引位置的交易
                tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
                if err != nil {
                        log.Fatal(err)
                }

                fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
        }

        txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
        tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
        fmt.Println(isPending)       // false
}