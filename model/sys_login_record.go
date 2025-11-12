package model

import (
	"bosh-admin/core/db"
)

type SysLoginRecord struct {
	db.BasicModel
	Uid          uint          `gorm:"comment:用户id" json:"uid"`
	Username     string        `gorm:"type:varchar(100);comment:用户名" json:"username"`
	LoginIP      string        `gorm:"column:login_ip;type:varchar(20);comment:登录ip" json:"loginIP"`
	LoginRegion  string        `gorm:"comment:登录地点" json:"loginRegion"`
	UserAgent    string        `gorm:"type:text;comment:用户代理" json:"userAgent"`
	LoginOS      string        `gorm:"column:login_os;comment:操作系统" json:"loginOS"`
	LoginBrowser string        `gorm:"comment:浏览器" json:"loginBrowser"`
	LoginStatus  int           `gorm:"default:0;comment:登录状态 1:成功 0:失败" json:"loginStatus"`
	LoginTime    db.CustomTime `gorm:"comment:登录时间" json:"loginTime"`
	LogoutTime   db.CustomTime `gorm:"comment:登出时间" json:"logoutTime"`
}

func (SysLoginRecord) TableName() string {
	return "sys_login_record"
}

func (SysLoginRecord) TableComment() string {
	return "系统登录记录表"
}
