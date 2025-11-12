package model

import (
	"bosh-admin/core/db"
)

// SysAppKey 应用API密钥表
type SysAppKey struct {
	db.BasicModel
	AppId     string `gorm:"type:varchar(100);not null;index;comment:应用ID" json:"appId"`
	AppKey    string `gorm:"type:varchar(100);not null;unique;comment:API密钥" json:"appKey"`
	SecretKey string `gorm:"type:varchar(100);not null;comment:密钥Secret" json:"secretKey"`
	Status    int    `gorm:"default:1;comment:状态 0禁用 1启用" json:"status"`
	ExpiredAt db.CustomTime `gorm:"comment:过期时间" json:"expiredAt"`
	Creator   string `gorm:"type:varchar(100);comment:创建者" json:"creator"`
	Updater   string `gorm:"type:varchar(100);comment:更新者" json:"updater"`
	Remark    string `gorm:"type:varchar(200);comment:备注" json:"remark"`
	
	// 关联应用信息
	App SysApp `gorm:"foreignKey:AppId;references:AppId" json:"app"`
}

func (SysAppKey) TableName() string {
	return "sys_app_key"
}

func (SysAppKey) TableComment() string {
	return "应用API密钥表"
}