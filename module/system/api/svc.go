package api

import (
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/model"
)

type SysApiSvc struct{}

func NewSysApiSvc() *SysApiSvc {
	return &SysApiSvc{}
}

func (svc *SysApiSvc) GetApiList(apiName, apiGroup, apiMethod, apiPath string, isRequired *int, pageNo, pageSize int) ([]model.SysApi, int64, error) {
	var list []model.SysApi
	var total int64
	var err error
	query := db.GormDB().Model(&model.SysApi{})
	if apiName != "" {
		query = query.Where("api_name LIKE ?", "%"+apiName+"%")
	}
	if apiGroup != "" {
		query = query.Where("api_group = ?", apiGroup)
	}
	if apiMethod != "" {
		query = query.Where("api_method = ?", apiMethod)
	}
	if apiPath != "" {
		query = query.Where("api_path LIKE ?", "%"+apiPath+"%")
	}
	if isRequired != nil {
		query = query.Where("is_required = ?", *isRequired)
	}
	if pageNo > 0 && pageSize > 0 {
		err = query.Count(&total).Error
		if err != nil {
			return nil, 0, exception.NewException("查询api数量失败", err)
		}
		query = query.Scopes(db.PageScope(pageNo, pageSize))
	}
	err = query.Find(&list).Error
	if err != nil {
		return nil, 0, exception.NewException("查询api列表失败", err)
	}
	return list, total, nil
}

func (svc *SysApiSvc) GetApiInfo(id uint) (*model.SysApi, error) {
	api, err := db.QueryById[model.SysApi](id)
	if err != nil {
		return nil, exception.NewException("查询api信息失败", err)
	}
	return api, nil
}

func (svc *SysApiSvc) AddApi(req AddApiReq) error {
	var count int64
	err := db.GormDB().Model(&model.SysApi{}).Where("api_method = ?").Where("api_path = ?", req.ApiPath).Count(&count).Error
	if err != nil {
		return exception.NewException("查询api信息失败", err)
	}
	if count > 0 {
		return exception.NewException("api已存在")
	}
	err = db.Create(&req, new(model.SysApi).TableName())
	if err != nil {
		return exception.NewException("新增api失败", err)
	}
	return nil
}

func (svc *SysApiSvc) EditApi(req EditApiReq) error {
	api, err := db.QueryById[model.SysApi](req.Id)
	if err != nil {
		return exception.NewException("查询api信息失败", err)
	}
	if api == nil {
		return exception.NewException("api不存在")
	}
	err = db.Updates(&req, new(model.SysApi).TableName())
	if err != nil {
		return exception.NewException("修改api失败", err)
	}
	return nil
}

func (svc *SysApiSvc) DelApi(id any) error {
	api, err := db.QueryById[model.SysApi](id)
	if err != nil {
		return exception.NewException("查询api信息失败", err)
	}
	if api == nil {
		return exception.NewException("api不存在")
	}
	err = db.DelById[model.SysApi](id)
	if err != nil {
		return exception.NewException("删除api失败", err)
	}
	return nil
}

func (svc *SysApiSvc) GetApiGroups() ([]string, error) {
	var groups []string
	err := db.GormDB().Model(&model.SysApi{}).Distinct("api_group").Pluck("api_group", &groups).Error
	if err != nil {
		return nil, exception.NewException("获取api分组失败", err)
	}
	return groups, nil
}
