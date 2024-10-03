# level2-memo

## 加密算法
- SHA3-256 和 Keccak-256
虽然 SHA3-256 和 Keccak-256 都是基于 Keccak 算法的,但由于标准化过程中的微小差异，它们在实际应用中被视为两个不同的哈希函数。
SHA3-256 通常用于需要严格遵循美国政府标准的场景，包括某些政府和军事应用。
Keccak-256 在以太坊等区块链技术中广泛使用，以太坊智能合约中的哈希函数使用的就是 Keccak-256，而非 SHA3-256。

## golang

go env GOPATH 
go env GOROOT

调用其他文件模块
go env
set GO111MODULE=auto

一、Golang推荐学习网站
https://gobyexample.com/



二、Gin框架
https://gin-gonic.com/zh-cn/docs/


```
solcjs --version  
```

client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
这里只是演示地址
尝试在上面四个BaaS平台注册账号后，创建endpoint，使用平台分配的地址

https://app.infura.io/  推荐
https://alchemy.com/
https://rockx.com/
https://quicknode.com/


## 区块浏览器

### Sepolia
Sepolia 是以太坊的一个较新的测试网络（testnet），设计用于替代一些旧的测试网络，如 Ropsten。Sepolia 测试网提供了一个环境，让开发者可以在生产环境（mainnet）部署之前，测试和调试他们的智能合约和应用程序。
主要特点
1. Proof of Authority (PoA)： Sepolia 使用了权威证明（PoA）共识机制。在 PoA 网络中，事务和区块的验证是由一组受信任的验证节点（通常被称为权威节点）执行的，这与更加去中心化的机制（如工作量证明，PoW）不同。这使得 Sepolia 能够提供更快的交易确认速度和更低的成本，因为它不依赖于能源密集型的挖矿过程。

2. 环境稳定性： 与老旧的测试网络相比，如 Ropsten，Sepolia 提供了更高的稳定性和可预测性，这对于开发者在测试和部署阶段尤其重要。

3. 开发和测试友好： Sepolia 设计为一个友好的开发和测试环境，使开发者能够轻松地部署和测试智能合约，执行交易并尝试各种区块链操作，而不用担心高成本或不稳定性。

使用 Sepolia 测试网络的优势
1. 无需真实资金： 开发者可以在 Sepolia 上执行交易和部署智能合约，而无需使用真实的以太币，这减少了开发过程中的风险和成本。

2. 接近主网操作： 尽管 Sepolia 使用的是 PoA 机制，它仍然提供了一个与以太坊主网相似的环境，帮助开发者为最终的主网部署做好准备。

3. 社区支持： 由于它是以太坊官方支持的测试网之一，开发者可以轻松获得文档支持和社区帮助。

- Faucet   
https://cloud.google.com/application/web3/faucet/ethereum/sepolia

- Browser
https://sepolia.etherscan.io/


## todolist
26_event_read_0xprotocol.go 可以顺利执行,没有触发LogFill的数据


solc solidity(0.8.0) require the same version
sudo npm install -g solc@latest   //0.8.27

## deploy
address 0x61774B7094Bc76bDf90c67b4f1aE66B00d1556B4
txt 0x4f31b244b0f2522e964c9564c9797577377d77aeff76ac09a819224bcb25f6c5