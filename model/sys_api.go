package model

import (
	"bosh-admin/core/db"
)

type SysApi struct {
	db.BasicModel
	ApiName    string `gorm:"type:varchar(128);not null;comment:api名称" json:"apiName"`
	ApiGroup   string `gorm:"type:varchar(128);not null;comment:api组" json:"apiGroup"`
	ApiPath    string `gorm:"type:varchar(128);not null;comment:api路径" json:"apiPath"`
	ApiMethod  string `gorm:"type:varchar(16);not null;comment:api方法" json:"apiMethod"`
	ApiDesc    string `gorm:"type:varchar(128);comment:api描述" json:"apiDesc"`
	IsRequired int    `gorm:"type:int;default:0;comment:是否必选 0否 1是" json:"isRequired"`
}

func (SysApi) TableName() string {
	return "sys_api"
}

func (SysApi) TableComment() string {
	return "系统api表"
}
