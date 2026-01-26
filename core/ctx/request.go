package ctx

import (
	"errors"
	"strings"

	"bosh-admin/global"

	"github.com/go-playground/validator/v10"
)

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// ValidateParams 校验请求参数
func (c *Context) ValidateParams(req any) (string, error) {
	err := c.ShouldBind(req)
	if err != nil {
		// 获取validator.ValidationErrors类型的errors
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			return ServerError, err
		}
		errsMap := removeTopStruct(errs.Translate(global.Trans))
		var errsArr []string
		for _, v := range errsMap {
			errsArr = append(errsArr, v)
		}
		return ParamsError, errors.New(strings.Join(errsArr, ";"))
	}
	return "", nil
}

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

// Sort 排序
type Sort struct {
	SortProp  string `json:"prop" form:"prop" validate:"omitempty"`                              // 排序字段
	SortOrder string `json:"order" form:"order" validate:"omitempty,oneof=ascending descending"` // 排序规则
}
