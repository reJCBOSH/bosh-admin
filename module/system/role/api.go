package role

import (
	"bosh-admin/core/ctx"
)

type SysRoleApi struct {
	svc *SysRoleSvc
}

func NewSysRoleApi() *SysRoleApi {
	return &SysRoleApi{
		svc: NewSysRoleSvc(),
	}
}

func (h *SysRoleApi) GetRoleList(c *ctx.Context) {
	var req GetRoleListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetRoleList(req.RoleName, req.RoleCode, req.Status, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithList(list, total)
}

func (h *SysRoleApi) GetRoleInfo(c *ctx.Context) {
	var req ctx.IdReq
	info, err := h.svc.GetRoleById(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(info)
}

func (h *SysRoleApi) AddRole(c *ctx.Context) {
	var req AddRoleReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.AddRole(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysRoleApi) EditRole(c *ctx.Context) {
	var req EditRoleReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.EditRole(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysRoleApi) DelRole(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelRole(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysRoleApi) GetRoleMenu(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	menus, err := h.svc.GetRoleMenu(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(menus)
}

func (h *SysRoleApi) GetRoleMenuIds(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	menuIds, err := h.svc.GetRoleMenuIds(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(menuIds)
}

func (h *SysRoleApi) SetRoleMenuAuth(c *ctx.Context) {
	var req SetRoleMenuAuthReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.SetRoleMenuAuth(req.RoleId, req.MenuIds)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysRoleApi) GetRoleDeptIds(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	deptIds, err := h.svc.GetRoleDeptIds(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(deptIds)
}

func (h *SysRoleApi) SetRoleDataPerm(c *ctx.Context) {
	var req SetRoleDataPermReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.SetRoleDataPerm(req.RoleId, req.DataPerm, req.DeptIds)
	if c.HandlerError(err) {
		return
	}
	// 判断是否统一角色
	userAuthInfo := c.GetUserAuthInfo()
	c.SuccessWithData(userAuthInfo.RoleId == req.RoleId)
}

func (h *SysRoleApi) SetRoleStatus(c *ctx.Context) {
	var req SetRoleStatusReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userAuthInfo := c.GetUserAuthInfo()
	err = h.svc.SetRoleStatus(userAuthInfo.RoleId, req.RoleId, req.Status)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}
