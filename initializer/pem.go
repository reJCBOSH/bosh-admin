package initializer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"bosh-admin/global"

	"github.com/duke-git/lancet/v2/fileutil"
)

func InitPem() {
	if !fileutil.IsExist(global.PrivateKeyFile) {
		if !fileutil.IsExist("keys") {
			err := fileutil.CreateDir("keys")
			if err != nil {
				panic(fmt.Sprintf("创建密钥目录失败: %v", err))
			}
		}
		// 生成私钥
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			panic(fmt.Sprintf("生成私钥失败: %v", err))
		}
		privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privateKeyBlock := pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		}
		privateKeyFile, err := os.Create(global.PrivateKeyFile)
		if err != nil {
			panic(fmt.Sprintf("创建私钥文件失败: %v", err))
		}
		err = pem.Encode(privateKeyFile, &privateKeyBlock)
		if err != nil {
			panic(fmt.Sprintf("编码私钥文件失败: %v", err))
		}
		_ = privateKeyFile.Close()

		// 生成公钥
		publicKey := privateKey.PublicKey
		publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
		if err != nil {
			panic(fmt.Sprintf("编码公钥文件失败: %v", err))
		}
		publicKeyBlock := pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		}
		publicKeyFile, err := os.Create(global.PublicKeyFile)
		if err != nil {
			panic(fmt.Sprintf("创建公钥文件失败: %v", err))
		}
		err = pem.Encode(publicKeyFile, &publicKeyBlock)
		if err != nil {
			panic(fmt.Sprintf("编码公钥文件失败: %v", err))
		}
		_ = publicKeyFile.Close()
	}
}
