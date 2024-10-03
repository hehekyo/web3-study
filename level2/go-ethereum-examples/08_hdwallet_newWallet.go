package main

import (
	"fmt"
	"log"

	"github.com/miguelmota/go-ethereum-hdwallet"
)

/**
使用了 go-ethereum-hdwallet 库来从一个助记词（mnemonic）生成以太坊钱包，并从该钱包派生出特定的账户地址。

功能和用途
HD钱包：这种钱包允许从单一的种子生成一系列的公钥和私钥对，增加了安全性和隐私性。
灵活性和安全性：通过使用不同的派生路径，可以为不同的目的生成多个地址，而不必为每个地址保留单独的密钥。


在加密货币领域中，"m/44'/60'/0'/0/0"是一个HD（层次确定性）钱包的派生路径，遵循了BIP44（Bitcoin Improvement Proposal 44）的标准。
这个路径定义了如何从一个单一的种子（master seed）派生出加密货币的私钥和公钥。下面我们来详细解析这个路径的每个部分的含义：

派生路径解析："m/44'/60'/0'/0/0"
m:

表示这是一个主节点（Master key）。
44':

这是BIP44的路径标识符，用于多币种和多账户钱包。
'（硬化）表示派生过程包括了额外的安全步骤，防止某些攻击，例如父私钥泄漏问题。
60':

这是以太坊的SLIP-0044币种代码。每种加密货币类型都有一个唯一的代码，比如比特币的代码是0'。
以太坊的代码为60'，表示这个路径是用来生成以太坊地址的。
0':

第一个账户编号。这允许用户为不同的目的创建多个账户，例如一个用于储蓄，另一个用于日常交易。
账户是隔离的，每个账户下可以生成多个地址。
0:

这表示地址在账户中的链类型。对于以太坊来说，通常使用0表示外部链（用于接收资金），1表示内部链（用于找零或交易变更）。
0:

这是此账户下的第一个地址索引。通过改变这个数字，可以生成同一账户下的更多地址。
使用场景
隐私性和安全性：通过使用不同的索引，可以生成多个地址，帮助提高隐私性，因为所有的交易不会都集中在同一个地址上。
组织性：通过使用不同的账户编号，可以将资金分散管理，便于不同用途或资金的隔离管理。
恢复性：只需记住初始的助记词和派生路径，就可以在任何支持BIP44的钱包软件中恢复出全部的子地址和相应的私钥。
**/

func main() {
        //生成钱包实例
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

        //解析派生路径并生成账户
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

        //生成第二个账户地址
	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}