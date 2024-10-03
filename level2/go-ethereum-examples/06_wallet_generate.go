package main

import (
        "crypto/ecdsa"
        "fmt"
        "log"

        "github.com/ethereum/go-ethereum/common/hexutil"
        "github.com/ethereum/go-ethereum/crypto"
        "golang.org/x/crypto/sha3"
)

func main() {
        // 1.生成私钥
        privateKey, err := crypto.GenerateKey()
        if err != nil {
                log.Fatal(err)
        }

        //2.私钥转换为字节并编码
        privateKeyBytes := crypto.FromECDSA(privateKey)
        fmt.Println("privateKeyBytes:",hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

        //3.获取公钥并进行类型断言
        publicKey := privateKey.Public()
        publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
        if !ok {
                log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
        }
        //4. 公钥转换为字节并编码
        publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
        fmt.Println("publicKeyBytes:",hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

        //5. 计算以太坊地址 这个函数计算公钥的 Keccak-256 哈希，然后取最后 20 字节作为地址。
        address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
        fmt.Println("ether address:",address) // 0x96216849c49358B10257cb55b28eA603c874b05E

        //手动计算地址的哈希
        //创建一个新的 Keccak-256 哈希函数的实例
        hash := sha3.NewLegacyKeccak256()
        //用于将公钥的字节数据（去除第一个字节，这个字节通常是用来标识公钥是压缩形式还是非压缩形式的前缀）添加到哈希函数中进行处理。
        //Keccak-256 生成的哈希值是32字节长
        hash.Write(publicKeyBytes[1:])
        //计算哈希值并取最后 20 字节，与以太坊地址计算方式一致。
        //hash.Sum(nil) Sum appends the current hash to b and returns the resulting slice.
        fmt.Println("hash:",hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}