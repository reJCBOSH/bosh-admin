package operationRecord

import (
	"bosh-admin/core/ctx"
)

type HandlerSysOperationRecord struct {
	svc *SvcSysOperationRecord
}

func NewHandlerSysOperationRecord() *HandlerSysOperationRecord {
	return &HandlerSysOperationRecord{
		svc: NewSvcSysOperationRecord(),
	}
}

func (h *HandlerSysOperationRecord) GetOperationRecordList(c *ctx.Context) {
	var req GetOperationRecordListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	records, total, err := h.svc.GetOperationRecordList(req.Username, req.Method, req.Path, req.RequestIP, req.Status, req.StartTime, req.EndTime, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	var list []OperationRecordListItem
	for _, record := range records {
		list = append(list, OperationRecordListItem{
			Id:             record.Id,
			CreatedAt:      record.CreatedAt.String(),
			Username:       record.Username,
			Method:         record.Method,
			Path:           record.Path,
			Status:         record.Status,
			Latency:        record.Latency,
			RequestIP:      record.RequestIP,
			RequestRegion:  record.RequestRegion,
			RequestOS:      record.RequestOS,
			RequestBrowser: record.RequestBrowser,
		})
	}
	c.SuccessWithList(list, total)
}

func (h *HandlerSysOperationRecord) GetOperationRecordInfo(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	data, err := h.svc.GetOperationRecordById(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(data)
}

func (h *HandlerSysOperationRecord) DelOperationRecord(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelOperationRecord(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysOperationRecord) BatchDelOperationRecord(c *ctx.Context) {
	var req ctx.IdsReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.BatchDelOperationRecord(req.Ids)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}
