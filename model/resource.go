package model

import (
	"bosh-admin/core/db"
)

type Resource struct {
	db.BasicModel
	Source         string `gorm:"comment:来源" json:"source"`
	IP             string `gorm:"column:ip;type:varchar(20);comment:IP" json:"IP"`
	FileName       string `gorm:"comment:文件名" json:"fileName"`
	FileSize       int64  `gorm:"comment:文件大小" json:"fileSize"`
	FileType       string `gorm:"comment:文件类型" json:"fileType"`
	MimeType       string `gorm:"comment:MIME类型" json:"mimeType"`
	StorePath      string `gorm:"comment:存储路径" json:"storePath"`
	FullPath       string `gorm:"comment:完整路径" json:"fullPath"`
	CheckSum       string `gorm:"type:varchar(32);comment:MD5校验和" json:"checkSum"`
	ReferenceCount int    `gorm:"comment:引用计数" json:"referenceCount"`
}

func (Resource) TableName() string {
	return "resource"
}

func (Resource) TableComment() string {
	return "资源表"
}
