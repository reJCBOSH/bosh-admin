package ctx

import (
	"bosh-admin/core/db"
	"bosh-admin/model"
)

type UserAuthInfo struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RoleId   uint   `json:"roleId"`
	RoleCode string `json:"roleCode"`
	DataPerm int    `json:"dataPerm"`
	DeptId   uint   `json:"deptId"`
	DeptCode string `json:"deptCode"`
	DeptPath string `json:"deptPath"`
}

func (c *Context) SetUserAuthInfo(info *model.SysUser) {
	userAuthInfo := &UserAuthInfo{
		UserId:   info.Id,
		Username: info.Username,
		Nickname: info.Nickname,
		RoleId:   info.RoleId,
		RoleCode: info.Role.RoleCode,
		DataPerm: info.Role.DataPerm,
		DeptId:   info.DeptId,
		DeptCode: info.Dept.DeptCode,
		DeptPath: info.Dept.DeptPath,
	}
	c.Set("userAuthInfo", userAuthInfo)
}

func (c *Context) GetUserAuthInfo() *UserAuthInfo {
	if info, exists := c.Get("userAuthInfo"); exists {
		return info.(*UserAuthInfo)
	} else {
		return nil
	}
}

func (c *Context) SetUserDataPerm(info *model.SysUser) {
	userDataPerm := &db.DataPermission{
		UserId:   info.Id,
		RoleId:   info.RoleId,
		RoleCode: info.Role.RoleCode,
		DataPerm: info.Role.DataPerm,
		DeptId:   info.DeptId,
		DeptPath: info.Dept.DeptPath,
	}
	c.Set("userDataPerm", userDataPerm)
}

func (c *Context) GetUserDataPerm() *db.DataPermission {
	if info, exists := c.Get("userDataPerm"); exists {
		return info.(*db.DataPermission)
	} else {
		return nil
	}
}
