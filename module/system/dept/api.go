package dept

import (
	"bosh-admin/core/ctx"
	"bosh-admin/module/basic"
)

type SysDeptApi struct {
	svc *SysDeptSvc
}

func NewSysDeptApi() *SysDeptApi {
	return &SysDeptApi{
		svc: NewSysDeptSvc(),
	}
}

func (h *SysDeptApi) GetDeptTree(c *ctx.Context) {
	list, err := h.svc.GetDeptTree()
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(list)
}

func (h *SysDeptApi) GetDeptList(c *ctx.Context) {
	var req GetDeptListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetDeptList(req.DeptName, req.DeptCode, req.Status, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithList(list, total)
}

func (h *SysDeptApi) GetDeptInfo(c *ctx.Context) {
	var req basic.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	info, err := h.svc.GetDeptById(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(info)
}

func (h *SysDeptApi) AddDept(c *ctx.Context) {
	var req AddDeptReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.AddDept(req)
	if c.HandlerError(err) {
		return
	}
	c.Success("添加成功")
}

func (h *SysDeptApi) EditDept(c *ctx.Context) {
	var req EditDeptReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.EditDept(req)
	if c.HandlerError(err) {
		return
	}
	c.Success("修改成功")
}

func (h *SysDeptApi) DelDept(c *ctx.Context) {
	var req basic.IdsReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelDept(req.Ids)
	if c.HandlerError(err) {
		return
	}
	c.Success("删除成功")
}
