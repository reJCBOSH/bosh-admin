package model

import (
	"bosh-admin/core/db"
)

// SysAppPerm 应用权限表
type SysAppPerm struct {
	db.BasicModel
	AppId     string `gorm:"type:varchar(100);not null;index;comment:应用ID" json:"appId"`
	Module    string `gorm:"type:varchar(100);not null;comment:模块名称" json:"module"`
	APIPath   string `gorm:"type:varchar(200);not null;comment:API路径" json:"apiPath"`
	Method    string `gorm:"type:varchar(10);not null;comment:请求方法" json:"method"`
	Status    int    `gorm:"default:1;comment:状态 0禁用 1启用" json:"status"`
	Creator   string `gorm:"type:varchar(100);comment:创建者" json:"creator"`
	Updater   string `gorm:"type:varchar(100);comment:更新者" json:"updater"`
	Remark    string `gorm:"type:varchar(200);comment:备注" json:"remark"`
	
	// 关联应用信息
	App SysApp `gorm:"foreignKey:AppId;references:AppId" json:"app"`
}

func (SysAppPerm) TableName() string {
	return "sys_app_perm"
}

func (SysAppPerm) TableComment() string {
	return "应用权限表"
}