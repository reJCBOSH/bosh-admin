package auth

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/log"
	"bosh-admin/domain/api/system/loginRecord"
	"bosh-admin/service/jwt"
)

type HandlerAuth struct {
	svc            *SvcAuth
	jwtSvc         *jwt.SvcJWT
	loginRecordSvc *loginRecord.SvcSysLoginRecord
}

func NewHandlerAuth() *HandlerAuth {
	return &HandlerAuth{
		svc:            NewSvcAuth(),
		jwtSvc:         jwt.NewSvcJWT(),
		loginRecordSvc: loginRecord.NewSvcSysLoginRecord(),
	}
}

func (h *HandlerAuth) UserLogin(c *ctx.Context) {
	var req UserLoginReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	loginIP := c.ClientIP()
	userAgent := c.Request.UserAgent()
	user, err := h.svc.UserLogin(req.Username, req.Password)
	if c.HandlerError(err) {
		if user != nil {
			if err = h.loginRecordSvc.AddLoginRecord(user.Id, req.Username, loginIP, userAgent, 0); err != nil {
				log.Error("添加登录记录失败:", err.Error())
			}
		}
		return
	}
	accessToken, refreshToken, expiresAt, err := h.jwtSvc.UserLogin(user)
	if c.HandlerError(err) {
		if err = h.loginRecordSvc.AddLoginRecord(user.Id, req.Username, loginIP, userAgent, 0); err != nil {
			log.Error("添加登录记录失败:", err.Error())
		}
		return
	}
	if err = h.loginRecordSvc.AddLoginRecord(user.Id, req.Username, loginIP, userAgent, 1); err != nil {
		log.Error("添加登录记录失败:", err.Error())
	}
	c.SuccessWithData(TokenResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	})
}

func (h *HandlerAuth) RefreshToken(c *ctx.Context) {
	var req RefreshTokenReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	accessToken, refreshToken, expiresAt, err := h.jwtSvc.RefreshToken(req.RefreshToken)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(TokenResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	})
}
