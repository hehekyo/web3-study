package main

import (
        "context"
        "crypto/ecdsa"
        "fmt"
        "log"
        "math/big"
        "golang.org/x/crypto/sha3"

        "github.com/ethereum/go-ethereum"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/common/hexutil"
        "github.com/ethereum/go-ethereum/core/types"
        "github.com/ethereum/go-ethereum/crypto"
        // "github.com/ethereum/go-ethereum/crypto/sha3"
        "github.com/ethereum/go-ethereum/ethclient"
)

/**
如何转移 ERC-20 代币。
ERC-20 代币是基于以太坊平台的一种标准化的代币类型，其中“ERC”代表以太坊请求评论（Ethereum Request for Comments），
而“20”是提案的编号。ERC-20 是一套定义了代币的行为和规则的接口标准，使得不同的代币能够在以太坊生态系统中以统一的方式交互。
这种标准化极大地促进了各种代币的交换、钱包的开发、交易所的集成，以及其他金融服务的实施。

ERC-20 代币包含以下几个核心方法和两个事件：

核心方法:

totalSupply: 这个函数返回代币的总供应量。
balanceOf: 返回某个地址（账户）的代币余额。
transfer: 允许代币持有者将代币发送到另一个地址。
transferFrom: 允许代币的持有者批准另一个地址从其账户转移一定数量的代币，通常用于实现代币的交易功能。
approve: 允许一个地址（称为委托人）来提取调用者账户中的代币，直到达到指定数量。
allowance: 返回一个地址（委托人）还被允许从另一个地址（所有者）中提取多少代币。
事件:

Transfer: 当代币从一个地址转移到另一个地址时触发。
Approval: 当一个地址被批准从另一个地址提取代币时触发。

应用实例
代币如 USDT（泰达币）、LINK（Chainlink）、DAI（MakerDAO的稳定币）
**/

func main() {
        // 转账可以通过下面链接查询
        // https://sepolia.etherscan.io/address/0x1485458Ba21df2371Aa01a914493B53fFf1784e4
        // client, err := ethclient.Dial("https://rinkeby.infura.io")
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

        /**
        在以太坊网络中，nonce 是一个重要的概念，特别是在交易处理中。Nonce 是一个由账户维护的计数器，
        用于确保每个交易只被处理一次。在技术层面上，nonce 是一个账户发出的交易数，
        其目的是防止交易重放攻击（即重复执行相同的交易）。

        Nonce 的作用和特点
         防止重放攻击：Nonce 确保了同一个交易不能被网络重复接受。如果有人试图重新提交相同的交易（即使是完全相同的内容和签名），
         只要 nonce 已经被网络接受过，该交易就会被认为是无效的。

         顺序执行：Nonce 还确保了来自同一账户的交易按照发送顺序被执行。这避免了因交易顺序错乱而可能导致的问题，
         例如，当一个账户的余额依赖于前一个交易的结果时。
        **/

        /**
        Nonce（序号）在以太坊网络中用于确保每个从特定账户发出的交易都是唯一的，并且按照发出的顺序执行。这是通过以下几个机制实现的：
        1. 唯一性
        每个以太坊账户都有一个关联的 nonce，这个 nonce 从0开始，并且每次发出交易时都会递增。
        Nonce 是每个交易的一部分，并且在网络上广播时会一并发送。网络节点在处理交易时会检查 nonce 的值：

        如果交易的 nonce 是账户当前 nonce 的下一个值，交易将被接受并处理。
        如果交易的 nonce 太低（即该 nonce 的交易已经被处理），交易将被拒绝，因为它可能是重放的旧交易。
        如果交易的 nonce 太高，交易将进入待处理队列，直到填补了缺失的 nonce 值。
        通过这种机制，以太坊确保了每个交易都是基于其发出顺序唯一处理的。

        2. 顺序性
        由于节点在处理交易时会严格按照 nonce 的顺序来执行，这确保了来自同一账户的交易不会被乱序处理。
        例如，如果一个账户的当前 nonce 是 100，那么只有 nonce 为 101 的交易才会被网络接受和处理。
        这意味着，无论网络状况如何，交易的执行顺序都是由发送者控制的。
        **/
        fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
        nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
        if err != nil {
                log.Fatal(err)
        }

        value := big.NewInt(0) // in wei (0 eth)
        gasPrice, err := client.SuggestGasPrice(context.Background())
        if err != nil {
                log.Fatal(err)
        }

        toAddress := common.HexToAddress("0x7EA4F504727E86d77D7eEB679e91c760D6ba667D")
        //代币合约地址
        tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

        transferFnSignature := []byte("transfer(address,uint256)")
        // hash := sha3.NewKeccak256()
        hash := sha3.NewLegacyKeccak256()
        
        hash.Write(transferFnSignature)
        methodID := hash.Sum(nil)[:4]
        fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

        paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
        fmt.Println("paddedAddress:",hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

        amount := new(big.Int)
        amount.SetString("1000000000000000000000", 10) // 1000 tokens
        paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
        fmt.Println("paddedAmount:",hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

        var data []byte
        data = append(data, methodID...) //将 methodID 中的每个 byte 添加到 data 切片中
        data = append(data, paddedAddress...)
        data = append(data, paddedAmount...)

        gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
                To:   &toAddress,
                Data: data,
        })
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("gasLimit:",gasLimit) // 23256

        /**
         在以太坊中创建交易时，交易的发送方地址（from 地址）不是直接作为交易参数传递的，而是从签名中推导出来的。
         这就是为什么在使用 types.NewTransaction() 创建交易时，你会看到没有直接指定发送者地址的参数。
         一旦交易被签名，以太坊网络的节点可以从交易的签名中恢复出公钥，并由公钥计算出交易的发送者地址。
         这个地址用来验证交易的 nonce 和账户余额，确保交易的有效性。
        **/
        tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

        chainID, err := client.NetworkID(context.Background())
        if err != nil {
                log.Fatal(err)
        }

        signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
        if err != nil {
                log.Fatal(err)
        }

        err = client.SendTransaction(context.Background(), signedTx)
        if err != nil {
                log.Fatal("SendTransaction:",err)
        }

        //https://sepolia.etherscan.io/tx/0xd5e6f556cde760f4e91f300de7940d1c305008d27759c4fd4ef45049a2ed23e3
        fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xd5e6f556cde760f4e91f300de7940d1c305008d27759c4fd4ef45049a2ed23e3
}