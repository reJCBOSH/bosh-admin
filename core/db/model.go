package db

import (
	"gorm.io/gorm"
)

// BasicModel 基础模型
type BasicModel struct {
	Id        uint           `gorm:"primaryKey" json:"id"`        // Id
	CreatedAt CustomTime     `gorm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt CustomTime     `gorm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`              // 删除时间
}

// AddBasicModel 新增基础模型
type AddBasicModel struct {
	CreatedAt CustomTime `gorm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt CustomTime `gorm:"updated_at" json:"updatedAt"` // 更新时间
}

// EditBasicModel 修改基础模型
type EditBasicModel struct {
	Id        uint       `gorm:"primaryKey" json:"id" form:"id" validate:"required,gt=0"` // Id
	UpdatedAt CustomTime `gorm:"updated_at" json:"updatedAt" form:"updatedAt"`            // 更新时间
}

// DataPermission 数据权限
type DataPermission struct {
	UserId   uint   // 用户id
	RoleId   uint   // 角色id
	RoleCode string // 角色编码
	DataPerm int    // 数据权限
	DeptId   uint   // 部门id
	DeptPath string // 部门路径
}

// PermissionModel 权限模型
type PermissionModel struct {
	CreatedBy uint `gorm:"created_by;default:0" json:"-"` // 创建人id
	UpdatedBy uint `gorm:"updated_by;default:0" json:"-"` // 更新人id
}

// AddPermModel 新增权限模型
type AddPermModel struct {
	CreatedBy uint `gorm:"created_by" json:"-"` // 创建人id
	UpdatedBy uint `gorm:"updated_by" json:"-"` // 更新人id
}

// EditPermModel 修改权限模型
type EditPermModel struct {
	UpdatedBy uint `gorm:"updated_by" json:"-"` // 更新人id
}
