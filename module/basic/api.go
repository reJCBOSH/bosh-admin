package basic

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/module/auth"
	"bosh-admin/util"

	"github.com/mojocn/base64Captcha"
)

type BasicApi struct {
	svc    *BasicSvc
	jwtSvc *auth.JWTSvc
}

func NewBasicApi() *BasicApi {
	return &BasicApi{
		svc:    NewBasicSvc(),
		jwtSvc: auth.NewJWTSvc(),
	}
}

func (h *BasicApi) Captcha(c *ctx.Context) {
	capConfig := global.Config.Captcha
	driverDigit := &base64Captcha.DriverDigit{
		Height:   capConfig.ImgHeight,
		Width:    capConfig.ImgWidth,
		Length:   capConfig.KeyLong,
		MaxSkew:  0.7,
		DotCount: 80,
	}
	id, b64s, answer, err := util.GenerateCaptcha("digit", util.DriverParam{DriverDigit: driverDigit})
	if c.HandlerError(err, "验证码获取失败") {
		return
	}
	if util.IsDev() {
		log.Debug(answer)
	}
	c.SuccessWithData(CaptchaResp{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: capConfig.KeyLong,
	})
}
