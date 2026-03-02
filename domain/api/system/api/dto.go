package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/db"
)

type GetApiListReq struct {
	ctx.Pagination
	ApiName    string `json:"apiName" form:"apiName"`
	ApiGroup   string `json:"apiGroup" form:"apiGroup"`
	ApiMethod  string `json:"apiMethod" form:"apiMethod"`
	ApiPath    string `json:"apiPath" form:"apiPath"`
	IsRequired *int   `json:"isRequired" form:"isRequired"`
}

type AddApiReq struct {
	db.AddBasicModel
	ApiName    string `json:"apiName" form:"apiName" validate:"required,min=1,max=100"`
	ApiGroup   string `json:"apiGroup" form:"apiGroup" validate:"required,min=1,max=100"`
	ApiMethod  string `json:"apiMethod" form:"apiMethod" validate:"required,oneof=GET POST PUT DELETE"`
	ApiPath    string `json:"apiPath" form:"apiPath" validate:"required,min=1,max=100"`
	ApiDesc    string `json:"apiDesc" form:"apiDesc"`
	IsRequired int    `json:"isRequired" form:"isRequired" validate:"oneof=0 1"`
}

type EditApiReq struct {
	db.EditBasicModel
	ApiName    string `json:"apiName" form:"apiName" validate:"required,min=1,max=100"`
	ApiGroup   string `json:"apiGroup" form:"apiGroup" validate:"required,min=1,max=100"`
	ApiDesc    string `json:"apiDesc" form:"apiDesc"`
	IsRequired int    `json:"isRequired" form:"isRequired" validate:"oneof=0 1"`
}
