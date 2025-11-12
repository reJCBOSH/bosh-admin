package model

import (
	"bosh-admin/core/db"
)

type SysDept struct {
	db.BasicModel
	ParentId     uint      `gorm:"default:0;comment:父id" json:"parentId"`
	DeptName     string    `gorm:"type:varchar(30);not null;comment:部门名称" json:"deptName"`
	DeptCode     string    `gorm:"type:varchar(30);not null;unique;comment:部门编码" json:"deptCode"`
	DeptPath     string    `gorm:"default:0;comment:部门路径" json:"deptPath"`
	Remark       string    `gorm:"comment:备注" json:"remark"`
	Status       int       `gorm:"default:0;comment:状态 0停用 1启用" json:"status"`
	DisplayOrder int       `gorm:"type:int;default:0;comment:显示顺序" json:"displayOrder"`
	Children     []SysDept `gorm:"-" json:"children"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}

func (SysDept) TableComment() string {
	return "系统部门表"
}
