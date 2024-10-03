package main

import (
        "context"
        "crypto/ecdsa"
        "fmt"
        "log"
        "math/big"

        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        "github.com/ethereum/go-ethereum/crypto"
        "github.com/ethereum/go-ethereum/ethclient"

        exchange "go-ethereum-examples/contracts/exchange"
)

func main() {
        client, err := ethclient.Dial("https://sepolia.infura.io/v3/a1439f90e4fe49b0a48be70256fe8af1")
        if err != nil {
                log.Fatal(err)
        }

        privateKey, err := crypto.HexToECDSA("1c7bd28c7b452b76277b82e3aaddb90fbdbacf80e4ba45a5b9548c4dbb18ff9e")
        if err != nil {
                log.Fatal(err)
        }

        publicKey := privateKey.Public()
        publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
        if !ok {
                log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
        }

        fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
        nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
        if err != nil {
                log.Fatal(err)
        }

        gasPrice, err := client.SuggestGasPrice(context.Background())
        if err != nil {
                log.Fatal(err)
        }

        auth := bind.NewKeyedTransactor(privateKey)
        auth.Nonce = big.NewInt(int64(nonce))
        auth.Value = big.NewInt(0)     // in wei
        auth.GasLimit = uint64(300000) // in units
        auth.GasPrice = gasPrice

        // input := "1.0"
        address, tx, instance, err := exchange.DeployExchange(auth, client)
        if err != nil {
                log.Fatal(err)
        }

        // https://sepolia.etherscan.io/tx/0x76a9e58c36c83a2415c681387c9ca5be6adab0f19a67f09eb38bae1a298b7ad1
        fmt.Println(address.Hex())   // 0xE517000c18532c6201529f878A9A1edFc846398A
        fmt.Println(tx.Hash().Hex()) // 0x76a9e58c36c83a2415c681387c9ca5be6adab0f19a67f09eb38bae1a298b7ad1

        _ = instance
}