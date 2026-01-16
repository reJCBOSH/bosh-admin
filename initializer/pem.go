package initializer

import (
	"fmt"
	"os"

	"bosh-admin/global"
	"bosh-admin/util"

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
		privatePem, publicPem, err := util.GenerateRsaPem(2048)
		if err != nil {
			panic(err)
		}
		privateKeyFile, err := os.Create(global.PrivateKeyFile)
		if err != nil {
			panic(fmt.Sprintf("创建私钥文件失败: %v", err))
		}
		_, err = privateKeyFile.Write(privatePem)
		if err != nil {
			panic(fmt.Sprintf("写入私钥文件失败: %v", err))
		}
		_ = privateKeyFile.Close()

		publicKeyFile, err := os.Create(global.PublicKeyFile)
		if err != nil {
			panic(fmt.Sprintf("创建公钥文件失败: %v", err))
		}
		_, err = publicKeyFile.Write(publicPem)
		if err != nil {
			panic(fmt.Sprintf("写入公钥文件失败: %v", err))
		}
		_ = publicKeyFile.Close()
	}
}
