package basic

type CaptchaResp struct {
	CaptchaId     string `json:"captchaId"`     // 验证码Id
	PicPath       string `json:"picPath"`       // 验证码图片
	CaptchaLength int    `json:"captchaLength"` // 验证码长度
}

type LoginReq struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Captcha   string `json:"captcha" validate:"required"`
	CaptchaId string `json:"captchaId" validate:"required"`
}
