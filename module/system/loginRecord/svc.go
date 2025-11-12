package loginRecord

import (
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/model"
	"bosh-admin/util"
	"time"

	ua "github.com/mssola/user_agent"
)

type SysLoginRecordSvc struct {
}

func NewSysLoginRecordSvc() *SysLoginRecordSvc {
	return &SysLoginRecordSvc{}
}

func (svc *SysLoginRecordSvc) GetLoginRecordList(username, startTime, endTime string, status *int, pageNo, pageSize int) ([]model.SysLoginRecord, int64, error) {
	var list []model.SysLoginRecord
	var total int64
	var err error
	query := db.GormDB().Model(&model.SysLoginRecord{})
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if startTime != "" && endTime != "" {
		query = query.Where("login_time BETWEEN ? AND ?", startTime, endTime)
	}
	if status != nil {
		query = query.Where("login_status = ?", *status)
	}
	if pageNo > 0 && pageSize > 0 {
		err = query.Count(&total).Error
		if err != nil {
			return nil, 0, exception.NewException("查询登录日志数量失败", err)
		}
		query = query.Scopes(db.PageScope(pageNo, pageSize))
	}
	err = query.Find(&list).Error
	if err != nil {
		return nil, 0, exception.NewException("查询登录日志列表失败", err)
	}
	return list, total, nil
}

func (svc *SysLoginRecordSvc) AddLoginRecord(uid uint, username, loginIP, userAgent string, loginStatus int) error {
	var record = model.SysLoginRecord{
		Uid:         uid,
		Username:    username,
		LoginIP:     loginIP,
		UserAgent:   userAgent,
		LoginStatus: loginStatus,
		LoginTime:   db.CustomTime(time.Now().Local()),
	}
	record.LoginRegion = util.IP2Region(loginIP)
	UA := ua.New(userAgent)
	record.LoginOS = UA.OS()
	record.LoginBrowser, _ = UA.Browser()
	if loginStatus == 0 {
		record.LogoutTime = db.CustomTime(time.Now().Local())
	}
	err := db.Create(&record)
	if err != nil {
		return exception.NewException("新增登录日志失败", err)
	}
	return nil
}

func (svc *SysLoginRecordSvc) DelLoginRecord(id uint) error {
	err := db.DelById[model.SysLoginRecord](id)
	if err != nil {
		return exception.NewException("删除登录日志失败", err)
	}
	return nil
}

func (svc *SysLoginRecordSvc) DelLoginRecordByIds(ids []uint) error {
	err := db.DelByIds[model.SysLoginRecord](ids)
	if err != nil {
		return exception.NewException("批量删除登录日志失败", err)
	}
	return nil
}
