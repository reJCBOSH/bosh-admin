package auth

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/model"
	"bosh-admin/util"
	"fmt"
)

type AuthSvc struct{}

func NewAuthSvc() *AuthSvc {
	return &AuthSvc{}
}

func (svc *AuthSvc) UserLogin(username, password string) (*model.SysUser, error) {
	var user model.SysUser
	err := db.GormDB().Where("username = ?", username).Preload("Role").Preload("Dept").First(&user).Error
	if err != nil {
		return nil, exception.NewException("账号或密码错误", err)
	}
	if user.Status == 0 {
		return nil, exception.NewException("账号已冻结, 请联系管理员")
	}
	if !util.BcryptCheck(password, user.Password) {
		if user.PwdRemainTime == 1 {
			err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", user.Id).UpdateColumns(map[string]interface{}{"pwd_remain_time": user.PwdRemainTime - 1, "status": 0}).Error
			if err != nil {
				return nil, exception.NewException(ctx.ServerError, err)
			}
		}
		err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", user.Id).Update("pwd_remain_time", user.PwdRemainTime-1).Error
		if err != nil {
			return nil, exception.NewException(ctx.ServerError, err)
		}
		return nil, exception.NewException(fmt.Sprintf("账号或密码错误, 剩余尝试次数: %d", user.PwdRemainTime))
	}
	if user.PwdRemainTime < 5 {
		err = db.GormDB().Model(&model.SysUser{}).Where("id = ?", user.Id).UpdateColumn("pwd_remain_time", 5).Error
		if err != nil {
			return nil, exception.NewException(ctx.ServerError, err)
		}
	}
	return &user, nil
}
