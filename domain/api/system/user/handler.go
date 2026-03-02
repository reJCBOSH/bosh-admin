package user

import (
	"bosh-admin/core/ctx"
)

type HandlerSysUser struct {
	svc *SvcSysUser
}

func NewHandlerSysUser() *HandlerSysUser {
	return &HandlerSysUser{
		svc: NewSvcSysUser(),
	}
}

func (h *HandlerSysUser) GetUserList(c *ctx.Context) {
	var req GetUserListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetUserList(req.Username, req.Nickname, req.Gender, req.Status, req.RoleId, req.DeptId, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	var listData []UserListItem
	for _, user := range list {
		listData = append(listData, UserListItem{
			Id:       user.Id,
			Username: user.Username,
			Avatar:   user.Avatar,
			Nickname: user.Nickname,
			Gender:   user.Gender,
			Status:   user.Status,
			RoleId:   user.RoleId,
			DeptId:   user.DeptId,
			RoleName: user.Role.RoleName,
			RoleCode: user.Role.RoleCode,
			DeptName: user.Dept.DeptName,
			DeptCode: user.Dept.DeptCode,
			Remark:   user.Remark,
		})
	}
	c.SuccessWithList(listData, total)
}

func (h *HandlerSysUser) GetUserInfo(c *ctx.Context) {
	var req ctx.IdReq
	info, err := h.svc.GetUserById(req.Id)
	if c.HandlerError(err) {
		return
	}
	var data = UserListItem{
		Id:       info.Id,
		Username: info.Username,
		Avatar:   info.Avatar,
		Nickname: info.Nickname,
		Gender:   info.Gender,
		Status:   info.Status,
		RoleId:   info.RoleId,
		DeptId:   info.DeptId,
		RoleName: info.Role.RoleName,
		RoleCode: info.Role.RoleCode,
		DeptName: info.Dept.DeptName,
		DeptCode: info.Dept.DeptCode,
		Remark:   info.Remark,
	}
	c.SuccessWithData(data)
}

func (h *HandlerSysUser) AddUser(c *ctx.Context) {
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

func (h *HandlerSysUser) EditUser(c *ctx.Context) {
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

func (h *HandlerSysUser) DelUser(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userAuthInfo := c.GetUserAuthInfo()
	err = h.svc.DelUser(userAuthInfo.UserId, req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysUser) ResetPassword(c *ctx.Context) {
	var req ctx.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userAuthInfo := c.GetUserAuthInfo()
	err = h.svc.ResetPassword(userAuthInfo.UserId, req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysUser) SetUserStatus(c *ctx.Context) {
	var req SetUserStatusReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userAuthInfo := c.GetUserAuthInfo()
	err = h.svc.SetUserStatus(userAuthInfo.UserId, req.Id, req.Status)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysUser) GetSelfInfo(c *ctx.Context) {
	userAuthInfo := c.GetUserAuthInfo()
	info, err := h.svc.GetUserById(userAuthInfo.UserId)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(SelfInfo{
		UserId:       info.Id,
		Username:     info.Username,
		Avatar:       info.Avatar,
		Nickname:     info.Nickname,
		Gender:       info.Gender,
		Birthday:     info.Birthday.String(),
		Email:        info.Email,
		Mobile:       info.Mobile,
		Introduce:    info.Introduce,
		PwdUpdatedAt: info.PwdUpdatedAt.String(),
	})
}

func (h *HandlerSysUser) EditSelfInfo(c *ctx.Context) {
	var req EditSelfInfoReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userAuthInfo := c.GetUserAuthInfo()
	err = h.svc.EditSelfInfo(userAuthInfo.UserId, req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *HandlerSysUser) EditSelfPassword(c *ctx.Context) {
	var req EditSelfPasswordReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	userAuthInfo := c.GetUserAuthInfo()
	err = h.svc.EditSelfPassword(userAuthInfo.UserId, req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}
