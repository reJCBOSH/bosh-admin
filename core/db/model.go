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

// Other 分页、排序
type Other struct {
	Limit   int    // 限制
	Offset  int    // 偏移
	OrderBy string // 排序
}

// PermissionModel 权限模型
type PermissionModel struct {
	CreatedBy uint   `gorm:"created_by" json:"-"` // 创建人id
	Creator   string `gorm:"creator" json:"-"`    // 创建人
	UpdatedBy uint   `gorm:"updated_by" json:"-"` // 更新人id
	Updater   string `gorm:"updater" json:"-"`    // 更新人
	DeletedBy uint   `gorm:"deleted_by" json:"-"` // 删除人id
	Deleter   string `gorm:"deleter" json:"-"`    // 删除人
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

// AddPermissionModel 新增权限模型
type AddPermissionModel struct {
	CreatedBy uint   `gorm:"created_by" json:"-"` // 创建人id
	Creator   string `gorm:"creator" json:"-"`    // 创建人
	UpdatedBy uint   `gorm:"updated_by" json:"-"` // 更新人id
	Updater   string `gorm:"updater" json:"-"`    // 更新人
}

// EditPermissionModel 修改权限模型
type EditPermissionModel struct {
	UpdatedBy uint   `gorm:"updated_by" json:"-"`    // 更新人id
	Updater   string `gorm:"updater" json:"updater"` // 更新人
}
