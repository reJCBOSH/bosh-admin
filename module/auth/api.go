package auth

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/log"
	"bosh-admin/module/system/loginRecord"
)

type AuthApi struct {
	svc            *AuthSvc
	jwtSvc         *JWTSvc
	loginRecordSvc *loginRecord.SysLoginRecordSvc
}

func NewAuthApi() *AuthApi {
	return &AuthApi{
		svc:            NewAuthSvc(),
		jwtSvc:         NewJWTSvc(),
		loginRecordSvc: loginRecord.NewSysLoginRecordSvc(),
	}
}

func (h *AuthApi) UserLogin(c *ctx.Context) {
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

func (h *AuthApi) RefreshToken(c *ctx.Context) {
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