package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-ethereum-examples/contracts/store"
)

/**
演示了如何连接到以太坊区块链，过滤并读取特定智能合约地址的事件日志，
并解析这些日志以获取有关事件的详细信息。代码还演示了如何对事件签名进行哈希处理以验证事件类型。
**/

func main() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/a1439f90e4fe49b0a48be70256fe8af1")
	if err != nil {
		log.Fatal(err)
	}

	//配置事件日志过滤器
	contractAddress := common.HexToAddress("0x61774B7094Bc76bDf90c67b4f1aE66B00d1556B4")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6773399),
		ToBlock:   big.NewInt(6773399),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//获取并解析事件日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//循环处理每个日志条目
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		//打印区块哈希
		fmt.Println(vLog.BlockHash.Hex()) // 0xe5ecf2d87884dbc0bd26e7dd6e76a27fb37b412a0104cb0859d2a82fda2f5541
		fmt.Println(vLog.BlockNumber)     // 6773399
		fmt.Println(vLog.TxHash.Hex())    // 0x1d821d2a48ad67f3240f427e8d17ee3348b2c1268678927a6485356df1f6100c

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}

		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
		        log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}
		
		/**
		若您的 solidity 事件包含 indexed 事件类型，那么它们将成为_主题_而不是日志的数据属性的一部分。
		在 solidity 中您最多只能有 4 个主题，但只有 3 个可索引的事件类型。第一个主题总是事件的签名。
		我们的示例合约不包含可索引的事件，但如果它确实包含，这是如何读取事件主题。
		**/

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	//计算和验证事件签名的哈希
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
