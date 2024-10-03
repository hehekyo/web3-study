package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "os"

        "github.com/ethereum/go-ethereum/accounts/keystore"
)
/**
keystore 是什么?
在以太坊生态系统中，keystore 是一种安全的文件格式，用于存储以太坊账户的私钥。
这种文件通过用户提供的密码进行加密，确保私钥的安全。
Keystore 文件使用户能够相对安全地在不同的应用程序或设备之间转移和备份他们的密钥。

使用场景
个人使用：个人用户可以通过Keystore文件安全地管理他们的以太坊私钥，特别是当他们需要在不同设备或钱包应用之间迁移私钥时。
开发者：在开发以太坊应用时，开发者可以使用 Keystore 文件来管理用户的账户和密钥，为应用内的交易签名提供一种安全的方式。
**/
func createKs() {
        // 创建一个新的 keystore
        //加密参数 StandardScryptN 和 StandardScryptP 用于定义密钥文件的加密强度。
        ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
        password := "secret"
        //创建新账户 使用密码secret来加密私钥文件
        account, err := ks.NewAccount(password)
        if err != nil {
                log.Fatal(err)
        }

        //输出账户地址
        fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

}

func importKs() {
        file := "./tmp/UTC--2024-09-23T07-08-22.887739000Z--092d23c174b27f9b9e994939a44ad7a7cb195deb"
        //初始化 KeyStore
        ks := keystore.NewKeyStore("./tmp1", keystore.StandardScryptN, keystore.StandardScryptP)
        //读取密钥文件
        jsonBytes, err := ioutil.ReadFile(file)
        if err != nil {
                log.Fatal(err)
        }

        password := "secret"
        //导入账户
        account, err := ks.Import(jsonBytes, password, password)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

        if err := os.Remove(file); err != nil {
                log.Fatal(err)
        }
}

func main() {
        createKs()
        // importKs()
}