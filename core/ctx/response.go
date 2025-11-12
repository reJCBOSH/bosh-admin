package ctx

import (
	"errors"
	"fmt"
	"net/http"

	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/global"
)

const (
	SUCCESS = 200
	FAIL    = 400
)

const (
	Success         = "操作成功"
	ServerError     = "系统服务错误"
	ParamsError     = "参数错误"
	RecordNotFound  = "记录不存在"
	TooManyRequests = "请求过于频繁"
)

// Response 响应结构体
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// ListData 列表数据
type ListData struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

// Response 响应
func (c *Context) Response(code int, data any, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// Success 成功响应
func (c *Context) Success(msg ...string) {
	var resMsg = Success
	if len(msg) > 0 {
		resMsg = msg[0]
	}
	c.Response(SUCCESS, nil, resMsg)
}

// SuccessWithData 成功数据响应
func (c *Context) SuccessWithData(data any) {
	c.Response(SUCCESS, data, Success)
}

// SuccessWithList 成功列表响应
func (c *Context) SuccessWithList(list any, total int64) {
	c.Response(SUCCESS, ListData{
		List:  list,
		Total: total,
	}, Success)
}

// SuccessWithDetail 成功详情响应
func (c *Context) SuccessWithDetail(data any, msg string) {
	c.Response(SUCCESS, data, msg)
}

// Fail 失败响应
func (c *Context) Fail(msg ...string) {
	var resMsg = ServerError
	if len(msg) > 0 {
		resMsg = msg[0]
	}
	c.Response(FAIL, nil, resMsg)
}

// HandlerError 错误处理
func (c *Context) HandlerError(err error, msg ...string) bool {
	if err != nil {
		var logErr = err.Error()
		if len(msg) > 0 {
			logErr = fmt.Sprintf("%s: %s", msg[0], err.Error())
			c.Fail(msg...)
		} else {
			if errors.Is(err, db.NotFound) {
				logErr = RecordNotFound
				c.Fail(RecordNotFound)
			} else {
				ex := new(exception.Exception)
				if errors.As(err, &ex) {
					logErr = ex.Error()
					exceptionErr := ex.GetError()
					if exceptionErr != nil {
						logErr += ": " + exceptionErr.Error()
					}
					c.Fail(err.Error())
				} else {
					c.Fail()
				}
			}
		}
		global.Logger.Error(logErr)
		return true
	}
	return false
}

// UnAuthorized 未授权响应
func (c *Context) UnAuthorized(msg string) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: FAIL,
		Data: nil,
		Msg:  msg,
	})
}

// TooManyRequests 请求过多响应
func (c *Context) TooManyRequests() {
	c.JSON(http.StatusTooManyRequests, Response{
		Code: FAIL,
		Data: nil,
		Msg:  TooManyRequests,
	})
}
