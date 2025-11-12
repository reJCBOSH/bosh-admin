package model

import (
	"bosh-admin/core/db"
)

type SysUser struct {
	db.BasicModel
	Username      string        `gorm:"type:varchar(100);not null;unique;comment:用户名" json:"username"`
	Password      string        `gorm:"not null;comment:密码" json:"-"`
	PwdUpdatedAt  db.CustomTime `gorm:"comment:密码修改时间" json:"pwdUpdatedAt"`
	PwdRemainTime int           `gorm:"default:5;comment:剩余尝试次数" json:"-"`
	Avatar        string        `gorm:"comment:头像" json:"avatar"`
	Nickname      string        `gorm:"type:varchar(100);not null;comment:昵称" json:"nickname"`
	Gender        int           `gorm:"default:0;comment:性别 0未知 1男 2女" json:"gender"`
	Birthday      db.CustomDate `gorm:"type:date;comment:生日" json:"birthday"`
	Email         string        `gorm:"comment:邮箱" json:"email"`
	Mobile        string        `gorm:"type:varchar(20);comment:联系方式" json:"mobile"`
	Introduce     string        `gorm:"type:varchar(200);comment:个人简介" json:"introduce"`
	Status        int           `gorm:"default:0;comment:状态 0冻结 1正常" json:"status"`
	RoleId        uint          `gorm:"comment:角色id" json:"roleId"`
	DeptId        uint          `gorm:"comment:部门id" json:"deptId"`
	Remark        string        `gorm:"comment:备注" json:"remark"`
	Role          SysRole       `gorm:"foreignKey:RoleId;references:Id" json:"role"`
	Dept          SysDept       `gorm:"foreignKey:DeptId;references:Id" json:"dept"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (SysUser) TableComment() string {
	return "系统用户表"
}
