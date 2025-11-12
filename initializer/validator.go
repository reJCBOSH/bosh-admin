package initializer

import (
	"fmt"
	"reflect"
	"strings"

	"bosh-admin/global"

	"github.com/dlclark/regexp2"
	lancetValidator "github.com/duke-git/lancet/v2/validator"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// InitValidator 初始化校验器
func InitValidator() {
	// 使用Validator引擎
	validate, _ := binding.Validator.Engine().(*validator.Validate)
	// 使用validate标签
	validate.SetTagName("validate")
	// 注册获取 json tag 的函数
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// 注册中文翻译器
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	global.Trans, _ = uni.GetTranslator("zh")
	err := zhTranslations.RegisterDefaultTranslations(validate, global.Trans)
	if err != nil {
		panic(fmt.Errorf("注册中文翻译器失败: %s \n", err.Error()))
	}

	// 注册自定义检验方法
	_ = validate.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		return lancetValidator.IsChineseMobile(fl.Field().String())
	})
	_ = validate.RegisterValidation("idnum", func(fl validator.FieldLevel) bool {
		return lancetValidator.IsChineseIdNum(fl.Field().String())
	})
	_ = validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return lancetValidator.IsChinesePhone(fl.Field().String())
	})
	_ = validate.RegisterValidation("pwd", func(fl validator.FieldLevel) bool {
		regex, _ := regexp2.Compile(`^(?=.*\d)(?=.*[A-Za-z])(?=.*[!@#$%^&*?\.])[A-Za-z0-9!@#$%^&*?\.]{8,16}$`, 0)
		result, _ := regex.MatchString(fl.Field().String())
		return result
	})
	_ = validate.RegisterValidation("creditcard", func(fl validator.FieldLevel) bool {
		return lancetValidator.IsCreditCard(fl.Field().String())
	})

	// 注册自定义错误提示
	_ = validate.RegisterTranslation("mobile", global.Trans, registerTranslator("mobile", "{0}格式不正确"), translate)
	_ = validate.RegisterTranslation("idnum", global.Trans, registerTranslator("idnum", "{0}格式不正确"), translate)
	_ = validate.RegisterTranslation("phone", global.Trans, registerTranslator("phone", "{0}格式不正确"), translate)
	_ = validate.RegisterTranslation("pwd", global.Trans, registerTranslator("pwd", "{0}格式不正确"), translate)
	_ = validate.RegisterTranslation("creditcard", global.Trans, registerTranslator("creditcard", "{0}格式不正确"), translate)
}

// registerTranslator 为自定义校验字段添加翻译功能
func registerTranslator(tag, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
