package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

/*
*
构建原始交易（Raw Transaction）
*
*/
func main() {
	// 连接到以太坊节点
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/a1439f90e4fe49b0a48be70256fe8af1")
	if err != nil {
		log.Fatal(err)
	}

	// 私钥
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

	value := big.NewInt(10000000000000000) // in wei (0.01 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x7EA4F504727E86d77D7eEB679e91c760D6ba667D")
	var data []byte
	// 创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// ts := types.Transactions{signedTx}
	// rawTxBytes := ts.GetRlp(0)
	// rawTxHex := hex.EncodeToString(rawTxBytes)

	// 使用 RLP 编码已签名的交易
	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatal(err)
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Printf("Encoded TX: %s\n", rawTxHex) // 打印 RLP 编码的交易
	//f86f05852993f5ae50825208947ea4f504727e86d77d7eeb679e91c760d6ba667d872386f26fc10000808401546d72a0a614a80fc0b0555d82320dc7a56e7fe398e4c29ab541049a504763af07314392a04374c85c29791df072fe4c8d6eea818320b2b6de471da79104b3588f6dcf3916
}
