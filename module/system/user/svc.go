package user

import (
	"time"

	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/global"
	"bosh-admin/model"
	"bosh-admin/util"
)

type SysUserSvc struct{}

func NewSysUserSvc() *SysUserSvc {
	return &SysUserSvc{}
}

func (svc *SysUserSvc) GetUserList(username, nickname string, gender, status *int, roleId, deptId *uint, pageNo, pageSize int) ([]model.SysUser, int64, error) {
	var list []model.SysUser
	var total int64
	var err error
	query := db.GormDB().Model(&model.SysUser{})
	if username != "" {
		query = query.Where("username Like ?", "%"+username+"%")
	}
	if nickname != "" {
		query = query.Where("nickname Like ?", "%"+nickname+"%")
	}
	if gender != nil {
		query = query.Where("gender = ?", *gender)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if roleId != nil {
		query = query.Where("role_id = ?", *roleId)
	}
	if deptId != nil {
		query = query.Where("dept_id = ?", *deptId)
	}
	if pageNo > 0 && pageSize > 0 {
		err = query.Count(&total).Error
		if err != nil {
			return nil, 0, exception.NewException("查询用户数量失败", err)
		}
		query = query.Scopes(db.PageScope(pageNo, pageSize))
	}
	err = query.Preload("Role").Preload("Dept").Find(&list).Error
	if err != nil {
		return nil, 0, exception.NewException("查询用户列表失败", err)
	}
	return list, total, nil
}

func (svc *SysUserSvc) GetUserById(id any) (*model.SysUser, error) {
	var user model.SysUser
	err := db.GormDB().Model(&model.SysUser{}).Where("id = ?", id).Preload("Role").Preload("Dept").First(&user).Error
	if err != nil {
		return nil, exception.NewException("查询用户失败", err)
	}
	return &user, nil
}

func (svc *SysUserSvc) AddUser(user AddUserReq) error {
	var count int64
	err := db.GormDB().Where("username = ?", user.Username).Count(&count).Error
	if err != nil {
		return exception.NewException("查询重名用户失败", err)
	}
	if count > 0 {
		return exception.NewException("用户名已存在")
	}
	user.Password, err = util.BcryptHash(user.Password)
	if err != nil {
		return exception.NewException("密码加密失败", err)
	}
	err = db.Create(&user, new(model.SysUser).TableName())
	if err != nil {
		return exception.NewException("新增用户失败", err)
	}
	return nil
}

func (svc *SysUserSvc) EditUser(user EditUserReq) error {
	originUser, err := db.QueryById[model.SysUser](user.Id)
	if err != nil {
		return exception.NewException("查询用户失败", err)
	}
	if user.Username != originUser.Username {
		var count int64
		err = db.GormDB().Where("username = ?", user.Username).Count(&count).Error
		if err != nil {
			return exception.NewException("查询重名用户失败", err)
		}
		if count > 0 {
			return exception.NewException("用户名已存在")
		}
	}
	err = db.Updates(&user, new(model.SysUser).TableName())
	if err != nil {
		return exception.NewException("修改用户失败", err)
	}
	return nil
}

func (svc *SysUserSvc) DelUser(currentUserId uint, id any) error {
	if currentUserId == id {
		return exception.NewException("不能删除自己")
	}
	var user model.SysUser
	err := db.GormDB().Where("id = ?", id).Preload("Role").First(&user).Error
	if err != nil {
		return exception.NewException("查询用户失败", err)
	}
	if user.Role.RoleCode == global.SuperAdmin {
		return exception.NewException("超级管理员用户不能删除")
	}
	err = db.DelById[model.SysUser](id)
	if err != nil {
		return exception.NewException("删除用户失败", err)
	}
	return nil
}

func (svc *SysUserSvc) ResetPassword(currentUserId uint, id any) error {
	if currentUserId == id {
		return exception.NewException("不能重置自己密码")
	}
	var user model.SysUser
	err := db.GormDB().Where("id = ?", id).Preload("Role").First(&user).Error
	if err != nil {
		return exception.NewException("查询用户失败", err)
	}
	if user.Role.RoleCode == global.SuperAdmin {
		return exception.NewException("超级管理员用户不能重置密码")
	}
	defaultPassword, err := util.BcryptHash(global.DefaultPassword)
	if err != nil {
		return exception.NewException("密码加密失败", err)
	}
	err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"password":        defaultPassword,
		"pwd_remain_time": 5,
		"pwd_updated_at":  time.Now(),
	}).Error
	if err != nil {
		return exception.NewException("重置密码失败", err)
	}
	return nil
}

func (svc *SysUserSvc) SetUserStatus(currentUserId uint, id any, status int) error {
	if currentUserId == id {
		return exception.NewException("不能修改自己状态")
	}
	var user model.SysUser
	err := db.GormDB().Where("id = ?", id).Preload("Role").First(&user).Error
	if err != nil {
		return exception.NewException("查询用户失败", err)
	}
	if user.Role.RoleCode == global.SuperAdmin {
		return exception.NewException("超级管理员用户不能修改状态")
	}
	if user.Status != status {
		err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", id).Update("status", status).Error
		if err != nil {
			return exception.NewException("修改用户状态失败", err)
		}
	}
	return nil
}

func (svc *SysUserSvc) EditSelfInfo(currentUserId uint, req EditSelfInfoReq) error {
	if currentUserId != req.Id {
		return exception.NewException("不能修改其他用户信息")
	}
	err := db.Updates(&req, new(model.SysUser).TableName())
	if err != nil {
		return exception.NewException("修改用户信息失败", err)
	}
	return nil
}

func (svc *SysUserSvc) EditSelfPassword(currentUserId uint, req EditSelfPasswordReq) error {
	user, err := db.QueryById[model.SysUser](currentUserId)
	if err != nil {
		return exception.NewException("查询用户失败", err)
	}
	if !util.BcryptCheck(req.OldPassword, user.Password) {
		return exception.NewException("旧密码错误")
	}
	newPassword, err := util.BcryptHash(req.NewPassword)
	if err != nil {
		return exception.NewException("密码加密失败", err)
	}
	err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", currentUserId).Updates(map[string]interface{}{
		"password":        newPassword,
		"pwd_remain_time": 5,
		"pwd_updated_at":  time.Now(),
	}).Error
	if err != nil {
		return exception.NewException("修改密码失败", err)
	}
	return nil
}
