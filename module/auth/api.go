package auth

import "bosh-admin/core/ctx"

type AuthApi struct {
	svc    *AuthSvc
	jwtSvc *JWTSvc
}

func NewAuthApi() *AuthApi {
	return &AuthApi{
		svc:    NewAuthSvc(),
		jwtSvc: NewJWTSvc(),
	}
}

func (h *AuthApi) UserLogin(c *ctx.Context) {
	var req UserLoginReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	user, err := h.svc.UserLogin(req.Username, req.Password)
	if c.HandlerError(err) {
		return
	}
	accessToken, refreshToken, expiresAt, err := h.jwtSvc.UserLogin(user)
	if c.HandlerError(err) {
		return
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
