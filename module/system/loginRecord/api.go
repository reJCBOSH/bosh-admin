package loginRecord

import (
	"bosh-admin/core/ctx"
	"bosh-admin/module/basic"
)

type SysLoginRecordApi struct {
	svc *SysLoginRecordSvc
}

func NewSysLoginRecordApi() *SysLoginRecordApi {
	return &SysLoginRecordApi{
		svc: NewSysLoginRecordSvc(),
	}
}

func (h *SysLoginRecordApi) GetLoginRecordList(c *ctx.Context) {
	var req GetLoginRecordListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetLoginRecordList(req.Username, req.StartTime, req.EndTime, req.Status, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithList(list, total)
}

func (h *SysLoginRecordApi) DelLoginRecord(c *ctx.Context) {
	var req basic.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelLoginRecord(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysLoginRecordApi) BatchDelLoginRecord(c *ctx.Context) {
	var req basic.IdsReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelLoginRecordByIds(req.Ids)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}
