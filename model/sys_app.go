package model

import (
	"bosh-admin/core/db"
)

// SysApp 外部应用表
type SysApp struct {
	db.BasicModel
	AppId     string `gorm:"type:varchar(100);not null;unique;comment:应用ID" json:"appId"`
	AppSecret string `gorm:"type:varchar(100);not null;comment:应用密钥" json:"appSecret"`
	AppName   string `gorm:"type:varchar(100);not null;comment:应用名称" json:"appName"`
	AppDesc   string `gorm:"type:text;comment:应用描述" json:"appDesc"`
	Status    int    `gorm:"default:1;comment:状态 0禁用 1启用" json:"status"`
	Creator   string `gorm:"type:varchar(100);comment:创建者" json:"creator"`
	Updater   string `gorm:"type:varchar(100);comment:更新者" json:"updater"`
	Remark    string `gorm:"type:varchar(200);comment:备注" json:"remark"`
}

func (SysApp) TableName() string {
	return "sys_app"
}

func (SysApp) TableComment() string {
	return "外部应用表"
}