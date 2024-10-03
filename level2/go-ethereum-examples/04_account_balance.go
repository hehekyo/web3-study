package main

import (
        "context"
        "fmt"
        "log"
        "math"
        "math/big"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
        client, err := ethclient.Dial("https://cloudflare-eth.com")
        if err != nil {
                log.Fatal(err)
        }

        account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")


        //client.BalanceAt() 获取给定地址在特定区块号下的余额。若区块号为 nil，则返回最新确认的区块的余额。
        //context.Background() 是 context 包中的一个函数，它返回一个空的 Context。
        //这个空的 Context 通常用作整个程序的主上下文，或者作为传递给其他函数的默认上下文。
        balance, err := client.BalanceAt(context.Background(), account, nil)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(balance) // 25893180161173005034

        //blockNumber 设置为 5532993 再次调用 client.BalanceAt()，但这次传入特定的区块号来获取那个时点的余额。
        blockNumber := big.NewInt(5532993)
        balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(balanceAt) // 25729324269165216042

        fbalance := new(big.Float)
        fbalance.SetString(balanceAt.String())

        //wei/10^18
        //.Quo(fbalance, divisor) 是一个方法调用，执行除法操作
        ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
        fmt.Println(ethValue) // 25.729324269165216041

        //待处理的余额
        pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
        fmt.Println(pendingBalance) // 25729324269165216042
}