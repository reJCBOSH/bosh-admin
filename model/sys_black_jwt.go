package model

import (
	"bosh-admin/core/db"
)

type SysBlackJwt struct {
	db.BasicModel
	BlackJwt  string `gorm:"comment:作废JWT" json:"blackJwt"`
	BlackUUID string `gorm:"black_uuid;comment:作废UUID" json:"blackUUID"`
}

func (SysBlackJwt) TableName() string {
	return "sys_black_jwt"
}

func (SysBlackJwt) TableComment() string {
	return "系统作废JWT表"
}
