package menu

import (
	"bosh-admin/core/ctx"
)

type HandlerSysMenu struct {
	svc *SvcSysMenu
}

func NewHandlerSysMenu() *HandlerSysMenu {
	return &HandlerSysMenu{
		svc: NewSvcSysMenu(),
	}
}

func (h *HandlerSysMenu) GetMenuTree(c *ctx.Context) {
	menu, err := h.svc.GetMenuTree()
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(menu)
}

func (h *HandlerSysMenu) GetMenuList(c *ctx.Context) {
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

func (h *HandlerSysMenu) GetMenuInfo(c *ctx.Context) {
	var req ctx.IdReq
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

func (h *HandlerSysMenu) AddMenu(c *ctx.Context) {
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

func (h *HandlerSysMenu) EditMenu(c *ctx.Context) {
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

func (h *HandlerSysMenu) DelMenu(c *ctx.Context) {
	var req ctx.IdReq
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

func (h *HandlerSysMenu) GetAsyncRoutes(c *ctx.Context) {
	userAuthInfo := c.GetUserAuthInfo()
	if userAuthInfo == nil {
		c.UnAuthorized("用户信息获取失败")
		return
	}
	routes, err := h.svc.GetAsyncRoutes(userAuthInfo.RoleId, userAuthInfo.RoleCode)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(routes)
}
