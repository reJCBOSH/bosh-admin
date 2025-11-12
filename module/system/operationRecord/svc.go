package operationRecord

import (
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/model"
)

type SysOperationRecordSvc struct{}

func NewSysOperationRecordSvc() *SysOperationRecordSvc {
	return &SysOperationRecordSvc{}
}

func (svc *SysOperationRecordSvc) GetOperationRecordList(username, method, path, requestIP string, status int, startTime, endTime string, pageNo, pageSize int) ([]model.SysOperationRecord, int64, error) {
	var list []model.SysOperationRecord
	var total int64
	var err error
	query := db.GormDB().Model(&model.SysOperationRecord{})
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if method != "" {
		query = query.Where("method = ?", method)
	}
	if path != "" {
		query = query.Where("path LIKE ?", "%"+path+"%")
	}
	if requestIp != "" {
		query = query.Where("request_ip LIKE ?", "%"+requestIp+"%")
	}
	if status != 0 {
		query = query.Where("status = ?", status)
	}
	if startTime != "" && endTime != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startTime, endTime)
	}
	if pageNo != 0 && pageSize != 0 {
		err = query.Count(&total).Error
		if err != nil {
			return nil, 0, exception.NewException("查询操作日志数量失败", err)
		}
		query = query.Scopes(db.PageScope(pageNo, pageSize))
	}
	err = query.Find(&list).Error
	if err != nil {
		return nil, 0, exception.NewException("查询操作日志列表失败", err)
	}
	return list, total, nil
}

func (svc *SysOperationRecordSvc) GetOperationRecordById(id uint) (*model.SysOperationRecord, error) {
	record, err := db.QueryById[model.SysOperationRecord](id)
	if err != nil {
		return nil, exception.NewException("查询操作日志失败", err)
	}
	return record, nil
}

func (svc *SysOperationRecordSvc) DelOperationRecord(id uint) error {
	err := db.DelById[model.SysOperationRecord](id)
	if err != nil {
		return exception.NewException("删除操作日志失败", err)
	}
	return nil
}

func (svc *SysOperationRecordSvc) BatchDelOperationRecord(ids []uint) error {
	err := db.DelByIds[model.SysOperationRecord](ids)
	if err != nil {
		return exception.NewException("批量删除操作日志失败", err)
	}
	return nil
}
