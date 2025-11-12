package util

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore

// DriverParam 验证码参数
type DriverParam struct {
	DriverAudio    *base64Captcha.DriverAudio
	DriverString   *base64Captcha.DriverString
	DriverChinese  *base64Captcha.DriverChinese
	DriverMath     *base64Captcha.DriverMath
	DriverLanguage *base64Captcha.DriverLanguage
	DriverDigit    *base64Captcha.DriverDigit
}

// GenerateCaptcha 生成验证码
func GenerateCaptcha(captchaType string, driverParam DriverParam) (id, b64s, answer string, err error) {
	var driver base64Captcha.Driver
	switch captchaType {
	case "audio":
		driver = driverParam.DriverAudio
	case "string":
		driver = driverParam.DriverString
	case "chinese":
		driver = driverParam.DriverChinese
	case "math":
		driver = driverParam.DriverMath
	case "language":
		driver = driverParam.DriverLanguage
	default:
		driver = driverParam.DriverDigit
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err = captcha.Generate()
	return
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id, answer string) bool {
	return store.Verify(id, answer, true)
}
