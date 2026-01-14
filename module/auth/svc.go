package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"

	"bosh-admin/core/ctx"
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/global"
	"bosh-admin/model"
	"bosh-admin/util"

	"github.com/pkg/errors"
)

type AuthSvc struct{}

func NewAuthSvc() *AuthSvc {
	return &AuthSvc{}
}

func (svc *AuthSvc) UserLogin(username, password string) (*model.SysUser, error) {
	privateKeyFile, err := os.Open(global.PrivateKeyFile)
	if err != nil {
		return nil, exception.NewException(ctx.ServerError, errors.Wrap(err, "打开私钥文件失败"))
	}
	privateKeyPEM, err := io.ReadAll(privateKeyFile)
	if err != nil {
		return nil, exception.NewException(ctx.ServerError, errors.Wrap(err, "读取私钥文件失败"))
	}
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, exception.NewException(ctx.ServerError, errors.New("私钥格式错误"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, exception.NewException(ctx.ServerError, errors.Wrap(err, "解析私钥失败"))
	}
	ciphertext, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return nil, exception.NewException(ctx.ServerError, errors.Wrap(err, "base64解密密码失败"))
	}
	decryptedPassword, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, exception.NewException(ctx.ServerError, errors.Wrap(err, "RSA解密密码失败"))
	}
	var user model.SysUser
	err = db.GormDB().Where("username = ?", username).Preload("Role").Preload("Dept").First(&user).Error
	if err != nil {
		return nil, exception.NewException("账号或密码错误", err)
	}
	if user.Status == 0 {
		return nil, exception.NewException("账号已冻结, 请联系管理员")
	}
	if !util.BcryptCheck(string(decryptedPassword), user.Password) {
		if user.PwdRemainTime == 1 {
			err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", user.Id).UpdateColumns(map[string]interface{}{"pwd_remain_time": user.PwdRemainTime - 1, "status": 0}).Error
			if err != nil {
				return nil, exception.NewException(ctx.ServerError, err)
			}
		}
		err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", user.Id).Update("pwd_remain_time", user.PwdRemainTime-1).Error
		if err != nil {
			return nil, exception.NewException(ctx.ServerError, err)
		}
		return nil, exception.NewException(fmt.Sprintf("账号或密码错误, 剩余尝试次数: %d", user.PwdRemainTime))
	}
	if user.PwdRemainTime < 5 {
		err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", user.Id).UpdateColumn("pwd_remain_time", 5).Error
		if err != nil {
			return nil, exception.NewException(ctx.ServerError, err)
		}
	}
	return &user, nil
}
