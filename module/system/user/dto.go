package user

import (
	"bosh-admin/core/db"
	"bosh-admin/module/basic"
)

type UserInfo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	RoleId   uint   `json:"roleId"`
	DeptId   uint   `json:"deptId"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
	DeptName string `json:"deptName"`
	DeptCode string `json:"deptCode"`
}

type GetUserListReq struct {
	basic.Pagination
	Username string `json:"username" form:"username"`                              // 用户名
	Nickname string `json:"nickname" form:"nickname"`                              // 昵称
	Gender   *int   `json:"gender" form:"gender" validate:"omitempty,oneof=0 1 2"` // 性别
	Status   *int   `json:"status" form:"status" validate:"omitempty,oneof=0 1"`   // 状态
	RoleId   *uint  `json:"roleId" form:"roleId" validate:"omitempty,gt=0"`        // 角色id
	DeptId   *uint  `json:"deptId" form:"deptId" validate:"omitempty,gt=0"`        // 部门id
}

type UserListItem struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"`
	Status   int    `json:"status"`
	RoleId   uint   `json:"roleId"`
	DeptId   uint   `json:"deptId"`
	Remark   string `json:"remark"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
	DeptName string `json:"deptName"`
	DeptCode string `json:"deptCode"`
}

type AddUserReq struct {
	db.AddBasicModel
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Gender   int    `json:"gender" validate:"oneof=0 1 2"`
	Status   int    `json:"status" validate:"oneof=0 1"`
	RoleId   uint   `json:"roleId" validate:"required,gt=0"`
	DeptId   uint   `json:"deptId" validate:"required,gt=0"`
	Remark   string `json:"remark"`
}

type EditUserReq struct {
	db.EditBasicModel
	Username string `json:"username" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Gender   int    `json:"gender" validate:"oneof=0 1 2"`
	Status   int    `json:"status" validate:"oneof=0 1"`
	RoleId   uint   `json:"roleId" validate:"required,gt=0"`
	DeptId   uint   `json:"deptId" validate:"required,gt=0"`
	Remark   string `json:"remark"`
}

type SetUserStatusReq struct {
	Id     uint `json:"id" validate:"required,gt=0"`
	Status int  `json:"status" validate:"oneof=0 1"`
}

type SelfInfo struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Gender    int    `json:"gender"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Introduce string `json:"introduce"`
}

type EditSelfInfoReq struct {
	db.EditBasicModel
	Avatar    string        `json:"avatar"`
	Nickname  string        `json:"nickname" validate:"required"`
	Gender    int           `json:"gender" validate:"oneof=0 1 2"`
	Birthday  db.CustomDate `json:"birthday"`
	Email     string        `json:"email"  validate:"omitempty,email"`
	Mobile    string        `json:"mobile" validate:"omitempty,mobile"`
	Introduce string        `json:"introduce"`
}

type EditSelfPasswordReq struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
	RePassword  string `json:"rePassword" validate:"required,eqfield=NewPassword"`
}
