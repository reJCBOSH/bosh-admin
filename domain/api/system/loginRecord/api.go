package loginRecord

import (
	"bosh-admin/core/ctx"
)

type HandlerSysLoginRecord struct {
	svc *SvcSysLoginRecord
}

func NewHandlerSysLoginRecord() *HandlerSysLoginRecord {
	return &HandlerSysLoginRecord{
		svc: NewSvcSysLoginRecord(),
	}
}

func (h *HandlerSysLoginRecord) GetLoginRecordList(c *ctx.Context) {
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

func (h *HandlerSysLoginRecord) DelLoginRecord(c *ctx.Context) {
	var req ctx.IdReq
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

func (h *HandlerSysLoginRecord) BatchDelLoginRecord(c *ctx.Context) {
	var req ctx.IdsReq
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
