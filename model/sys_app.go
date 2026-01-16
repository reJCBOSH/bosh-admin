package model

import (
	"bosh-admin/core/db"
)

// SysApp 外部应用表
type SysApp struct {
	db.BasicModel
	AppId         string `gorm:"type:varchar(100);not null;unique;comment:应用ID" json:"appId"`
	AppName       string `gorm:"type:varchar(100);not null;comment:应用名称" json:"appName"`
	AppDesc       string `gorm:"type:varchar(200);comment:应用描述" json:"appDesc"`
	Status        int    `gorm:"type:tinyint;default:1;comment:状态 0禁用 1启用" json:"status"`
	Remark        string `gorm:"type:varchar(200);comment:备注" json:"remark"`
	AppKey        string `gorm:"type:varchar(100);not null;unique;comment:应用密钥" json:"appKey"`
	SecretKey     string `gorm:"type:varchar(100);not null;comment:密钥Secret" json:"secretKey"`
	RsaPublicKey  string `gorm:"type:varchar(4000);comment:RSA公钥" json:"rsaPublicKey"`
	RsaPrivateKey string `gorm:"type:varchar(4000);comment:RSA私钥" json:"rsaPrivateKey"`
}

func (SysApp) TableName() string {
	return "sys_app"
}

func (SysApp) TableComment() string {
	return "外部应用表"
}
