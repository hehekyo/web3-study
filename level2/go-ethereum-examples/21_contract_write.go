package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-ethereum-examples/contracts/store"
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

        address := common.HexToAddress("0x61774B7094Bc76bDf90c67b4f1aE66B00d1556B4")
        instance, err := store.NewStore(address, client)
        if err != nil {
                log.Fatal(err)
        }

        key := [32]byte{}
        value := [32]byte{}
        copy(key[:], []byte("foo"))
        copy(value[:], []byte("bar"))

        tx, err := instance.SetItem(auth, key, value)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x1d821d2a48ad67f3240f427e8d17ee3348b2c1268678927a6485356df1f6100c

        result, err := instance.Items(nil, key)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(string(result[:])) // "bar"
}