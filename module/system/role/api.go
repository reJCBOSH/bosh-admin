package role

import (
	"bosh-admin/core/ctx"
	"bosh-admin/module/auth"
	"bosh-admin/module/basic"
)

type SysRoleApi struct {
	svc    *SysRoleSvc
	jwtSvc *auth.JWTSvc
}

func NewSysRoleApi() *SysRoleApi {
	return &SysRoleApi{
		svc:    NewSysRoleSvc(),
		jwtSvc: auth.NewJWTSvc(),
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
	var req basic.IdReq
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
	var req basic.IdReq
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
	var req basic.IdReq
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
	var req basic.IdReq
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
	var req basic.IdReq
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

func (h *SysRoleApi) SetRoleDataAuth(c *ctx.Context) {
	var req SetRoleDataAuthReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.SetRoleDataAuth(req.RoleId, req.DataAuth, req.DeptIds)
	if c.HandlerError(err) {
		return
	}
	// 判断是否统一角色
	userClaims := h.jwtSvc.GetUserClaims(c)
	c.SuccessWithData(userClaims.RoleId == req.RoleId)
}

func (h *SysRoleApi) SetRoleStatus(c *ctx.Context) {
	var req SetRoleStatusReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userCliams := h.jwtSvc.GetUserClaims(c)
	err = h.svc.SetRoleStatus(userCliams.RoleId, req.RoleId, req.Status)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}
