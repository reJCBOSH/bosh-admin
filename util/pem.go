package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

func GenerateRsaPem(bits int) (privatePem, publicPem []byte, err error) {
	// 生成私钥
	var privateKey *rsa.PrivateKey
	privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, errors.Wrap(err, "生成密钥对失败")
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem = pem.EncodeToMemory(&privateBlock)
	if privatePem == nil {
		return nil, nil, errors.New("生成私钥失败")
	}
	// 生成公钥
	publicKey := privateKey.PublicKey
	var publicKeyBytes []byte
	publicKeyBytes, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, nil, errors.Wrap(err, "编码公钥失败")
	}
	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem = pem.EncodeToMemory(&publicBlock)
	if publicPem == nil {
		return nil, nil, errors.New("生成公钥失败")
	}
	return privatePem, publicPem, nil
}

func RsaDecrypt(encryptedData []byte, privatePem []byte) ([]byte, error) {
	block, _ := pem.Decode(privatePem)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("私钥格式错误")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "解析私钥失败")
	}
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		return nil, errors.Wrap(err, "解密失败")
	}
	return decryptedData, nil
}
