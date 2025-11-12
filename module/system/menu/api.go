package menu

import (
	"bosh-admin/core/ctx"
	"bosh-admin/module/auth"
	"bosh-admin/module/basic"
)

type SysMenuApi struct {
	svc    *SysMenuSvc
	jwtSvc *auth.JWTSvc
}

func NewSysMenuApi() *SysMenuApi {
	return &SysMenuApi{
		svc:    NewSysMenuSvc(),
		jwtSvc: auth.NewJWTSvc(),
	}
}

func (h *SysMenuApi) GetMenuTree(c *ctx.Context) {
	menu, err := h.svc.GetMenuTree()
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(menu)
}

func (h *SysMenuApi) GetMenuList(c *ctx.Context) {
	var req GetMenuListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetMenuList(req.Title, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithList(list, total)
}

func (h *SysMenuApi) GetMenuInfo(c *ctx.Context) {
	var req basic.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	menu, err := h.svc.GetMenuById(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(menu)
}

func (h *SysMenuApi) AddMenu(c *ctx.Context) {
	var req AddMenuReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.AddMenu(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysMenuApi) EditMenu(c *ctx.Context) {
	var req EditMenuReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.EditMenu(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysMenuApi) DelMenu(c *ctx.Context) {
	var req basic.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelMenu(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysMenuApi) GetAsyncRoutes(c *ctx.Context) {
	userClaims := h.jwtSvc.GetUserClaims(c)
	if userClaims == nil {
		c.UnAuthorized("用户信息获取失败")
		return
	}
	routes, err := h.svc.GetAsyncRoutes(userClaims.RoleId, userClaims.RoleCode)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(routes)
}
