package role

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/db"
)

type GetRoleListReq struct {
	ctx.Pagination
	RoleName string `json:"roleName" form:"roleName"`
	RoleCode string `json:"roleCode" form:"roleCode"`
	Status   *int   `json:"status" form:"status" validate:"omitempty,oneof=0 1"`
}

type AddRoleReq struct {
	db.AddBasicModel
	RoleName string `json:"roleName" form:"roleName" validate:"required"`
	RoleCode string `json:"roleCode" form:"roleCode" validate:"required"`
	Remark   string `json:"remark" form:"remark"`
}

type EditRoleReq struct {
	db.EditBasicModel
	RoleName string `json:"roleName" form:"roleName" validate:"required"`
	Remark   string `json:"remark" form:"remark"`
}

type SetRoleMenuAuthReq struct {
	RoleId  uint   `json:"roleId" form:"roleId" validate:"required,gt=0"`
	MenuIds []uint `json:"menuIds" form:"menuIds" validate:"gt=0"`
}

type SetRoleDataPermReq struct {
	RoleId   uint   `json:"roleId" form:"roleId" validate:"required,gt=0"`
	DataPerm int    `json:"dataPerm" form:"dataPerm" validate:"required,oneof=1 2 3 4 5"`
	DeptIds  []uint `json:"deptIds" form:"deptIds"`
}

type SetRoleStatusReq struct {
	RoleId uint `json:"roleId" form:"roleId" validate:"required,gt=0"`
	Status int  `json:"status" form:"status" validate:"oneof=0 1"`
}
