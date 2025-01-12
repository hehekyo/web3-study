package main

import (
        "fmt"
        "log"
        "math"
        "math/big"
        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/ethclient"

        token "go-ethereum-examples/contracts/erc20_1"
)

func main() {
        client, err := ethclient.Dial("https://cloudflare-eth.com")
        if err != nil {
                log.Fatal(err)
        }

        
        // tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d") // Golem (GNT) Address
        tokenAddress := common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA")
        instance, err := token.NewToken(tokenAddress, client)
        if err != nil {
                log.Fatal(err)
        }

        address := common.HexToAddress("0xA9D1e08C7793af67e9d92fe308d5697FB81d3E43")
        bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
        if err != nil {
                log.Fatal(err)
        }

        name, err := instance.Name(&bind.CallOpts{})
        if err != nil {
                log.Fatal(err)
        }

        symbol, err := instance.Symbol(&bind.CallOpts{})
        if err != nil {
                log.Fatal(err)
        }

        decimals, err := instance.Decimals(&bind.CallOpts{})
        if err != nil {
                log.Fatal(err)
        }

        fmt.Printf("name: %s\n", name)         // "name: Golem Network"
        fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
        fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

        fmt.Printf("wei: %s\n", bal) // "wei: 271131065602071305846817"

        fbal := new(big.Float)
        fbal.SetString(bal.String())
        value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

        fmt.Printf("balance: %f", value) // "balance: 271131.065602"
}