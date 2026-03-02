package api

import "bosh-admin/core/ctx"

type HandlerSysApi struct {
	svc *SvcSysApi
}

func NewHandlerSysApi() *HandlerSysApi {
	return &HandlerSysApi{
		svc: NewSvcSysApi(),
	}
}

func (h *HandlerSysApi) GetApiList(c *ctx.Context) {
	var req GetApiListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetApiList(req.ApiName, req.ApiGroup, req.ApiMethod, req.ApiPath, req.IsRequired, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithList(list, total)
}

func (h *HandlerSysApi) GetApiInfo(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	data, err := h.svc.GetApiInfo(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(data)
}

func (h *HandlerSysApi) AddApi(c *ctx.Context) {
	var req AddApiReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.AddApi(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysApi) EditApi(c *ctx.Context) {
	var req EditApiReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.EditApi(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysApi) DelApi(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelApi(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysApi) GetApiGroups(c *ctx.Context) {
	groups, err := h.svc.GetApiGroups()
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(groups)
}
