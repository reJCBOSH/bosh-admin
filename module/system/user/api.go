package user

import (
	"bosh-admin/core/ctx"
	"bosh-admin/module/auth"
	"bosh-admin/module/basic"
)

type SysUserApi struct {
	svc    *SysUserSvc
	jwtSvc *auth.JWTSvc
}

func NewSysUserApi() *SysUserApi {
	return &SysUserApi{
		svc:    NewSysUserSvc(),
		jwtSvc: auth.NewJWTSvc(),
	}
}

func (h *SysUserApi) GetUserList(c *ctx.Context) {
	var req GetUserListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetUserList(req.Username, req.Nickname, req.Gender, req.Status, req.RoleId, req.DeptId, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithList(list, total)
}

func (h *SysUserApi) GetUserInfo(c *ctx.Context) {
	var req basic.IdReq
	info, err := h.svc.GetUserById(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(info)
}

func (h *SysUserApi) AddUser(c *ctx.Context) {
	var req AddUserReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.AddUser(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUserApi) EditUser(c *ctx.Context) {
	var req EditUserReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.EditUser(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUserApi) DelUser(c *ctx.Context) {
	var req basic.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userClaims := h.jwtSvc.GetUserClaims(c)
	err = h.svc.DelUser(userClaims.UserId, req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUserApi) ResetPassword(c *ctx.Context) {
	var req basic.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userClaims := h.jwtSvc.GetUserClaims(c)
	err = h.svc.ResetPassword(userClaims.UserId, req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUserApi) SetUserStatus(c *ctx.Context) {
	var req SetUserStatusReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userClaims := h.jwtSvc.GetUserClaims(c)
	err = h.svc.SetUserStatus(userClaims.UserId, req.Id, req.Status)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUserApi) EditSelfInfo(c *ctx.Context) {
	var req EditSelfInfoReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userClaims := h.jwtSvc.GetUserClaims(c)
	err = h.svc.EditSelfInfo(userClaims.UserId, req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysUserApi) EditSelfPassword(c *ctx.Context) {
	var req EditSelfPasswordReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userClaims := h.jwtSvc.GetUserClaims(c)
	err = h.svc.EditSelfPassword(userClaims.UserId, req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}
