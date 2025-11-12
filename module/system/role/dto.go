package role

import (
	"bosh-admin/core/db"
	"bosh-admin/module/basic"
)

type GetRoleListReq struct {
	basic.Pagination
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

type SetRoleDataAuthReq struct {
	RoleId   uint   `json:"roleId" form:"roleId" validate:"required,gt=0"`
	DataAuth int    `json:"dataAuth" form:"dataAuth" validate:"required,oneof=1 2 3 4 5"`
	DeptIds  []uint `json:"deptIds" form:"deptIds"`
}

type SetRoleStatusReq struct {
	RoleId uint `json:"roleId" form:"roleId" validate:"required,gt=0"`
	Status int  `json:"status" form:"status" validate:"oneof=0 1"`
}
