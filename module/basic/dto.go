package basic

// IdReq id请求
type IdReq struct {
	Id uint `json:"id" form:"id" validate:"required,min=1"` // id
}

// IdsReq ids请求
type IdsReq struct {
	Ids []uint `json:"ids" form:"ids" validate:"required,gt=0,dive,min=1"` // ids
}

// Pagination 分页
type Pagination struct {
	PageNo   int `json:"pageNo" form:"pageNo" validate:"required,min=-1,ne=0"`                       // 页码
	PageSize int `json:"pageSize" form:"pageSize" validate:"required_unless=PageNo -1|gt=0,max=100"` // 每页数量
}

// OrderBy 排序
type OrderBy struct {
	Field string `json:"field" form:"field" validate:"omitempty"`              // 排序字段
	Rule  string `json:"rule" form:"rule" validate:"omitempty,oneof=ASC DESC"` // 排序规则
}

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
