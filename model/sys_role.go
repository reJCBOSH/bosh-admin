package model

import (
	"bosh-admin/core/db"
)

type SysRole struct {
	db.BasicModel
	RoleName      string `gorm:"type:varchar(30);not null;comment:角色名称" json:"roleName" `
	RoleCode      string `gorm:"type:varchar(30);not null;unique;comment:角色编码" json:"roleCode" `
	Status        int    `gorm:"default:0;comment:状态 0冻结 1正常" json:"status" `
	Remark        string `gorm:"comment:备注" json:"remark" `
	DefaultRouter string `json:"comment:默认路由" json:"defaultRouter" `
	DataAuth      int    `gorm:"default:0;comment:数据权限 0未配置 1全部数据 2本部门数据 3本部门及以下数据 4本人数据 5自定义数据" json:"dataAuth" `
}

func (SysRole) TableName() string {
	return "sys_role"
}

func (SysRole) TableComment() string {
	return "系统角色表"
}
