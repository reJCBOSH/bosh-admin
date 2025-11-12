package model

import (
	"bosh-admin/core/db"
)

type SysOperationRecord struct {
	db.BasicModel
	Uid            uint   `gorm:"comment:用户id" json:"uid"`
	Username       string `gorm:"type:varchar(100);comment:用户名" json:"username"`
	Method         string `gorm:"comment:请求方式" json:"method"`
	Path           string `gorm:"comment:请求路径" json:"path"`
	Status         int    `gorm:"comment:请求状态" json:"status"`
	Latency        int64  `gorm:"comment:延迟" json:"latency"`
	UserAgent      string `gorm:"comment:用户代理" json:"userAgent"`
	RequestIP      string `gorm:"column:request_ip;type:varchar(20);comment:请求ip" json:"requestIP"`
	RequestRegion  string `gorm:"comment:请求地点" json:"requestRegion"`
	RequestOS      string `gorm:"column:request_os;comment:操作系统" json:"requestOS"`
	RequestBrowser string `gorm:"comment:浏览器" json:"requestBrowser"`
	RequestHeader  string `gorm:"type:text;comment:请求Header" json:"requestHeader"`
	RequestBody    string `gorm:"type:text;comment:请求Body" json:"requestBody"`
	ResponseHeader string `gorm:"type:text;comment:响应Header" json:"responseHeader"`
	ResponseBody   string `gorm:"type:text;comment:响应Body" json:"responseBody"`
}

func (SysOperationRecord) TableName() string {
	return "sys_operation_record"
}

func (SysOperationRecord) TableComment() string {
	return "系统操作记录表"
}
