package middleware

import (
	"bytes"
	"io"
	"strings"
	"time"

	"bosh-admin/core/ctx"
	"bosh-admin/module/openapi"
)

// OpenAPIAuth 开放API鉴权中间件
func OpenAPIAuth() ctx.HandlerFunc {
	svc := openapi.NewOpenAPISvc()

	return func(c *ctx.Context) {
		// 从请求头获取AppKey
		appKey := c.Request.Header.Get("X-App-Key")
		if appKey == "" {
			c.UnAuthorized("缺少AppKey")
			return
		}

		// 从请求头获取时间戳
		timestamp := c.Request.Header.Get("X-Timestamp")
		if timestamp == "" {
			c.UnAuthorized("缺少时间戳")
			return
		}

		// 验证时间戳有效性（防止重放攻击，允许5分钟的时间差）
		ts, err := time.Parse(time.RFC3339, timestamp)
		if err != nil {
			c.UnAuthorized("时间戳格式错误")
			return
		}

		if time.Since(ts).Abs() > 5*time.Minute {
			c.UnAuthorized("请求已过期")
			return
		}

		// 从请求头获取随机数
		nonce := c.Request.Header.Get("X-Nonce")
		if nonce == "" {
			c.UnAuthorized("缺少随机数")
			return
		}

		// 从请求头获取签名
		signature := c.Request.Header.Get("X-Signature")
		if signature == "" {
			c.UnAuthorized("缺少签名")
			return
		}

		// 获取请求体
		body, _ := c.GetRawData()
		// 重新设置请求体，以便后续处理可以读取
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// 验证签名
		err = svc.VerifySignature(appKey, timestamp, nonce, string(body), signature)
		if err != nil {
			c.UnAuthorized(err.Error())
			return
		}

		// 获取应用信息
		appKeyInfo, err := svc.GetAppKeyInfo(appKey)
		if err != nil {
			c.UnAuthorized(err.Error())
			return
		}

		// 将应用ID存储到上下文中，供后续使用
		c.Set("appId", appKeyInfo.AppId)

		// 验证通过，继续处理请求
		c.Next()
	}
}

// OpenAPIPermission 权限验证中间件
func OpenAPIPermission() ctx.HandlerFunc {
	svc := openapi.NewOpenAPISvc()

	return func(c *ctx.Context) {
		// 从上下文中获取应用ID
		appIdVal, exists := c.Get("appId")
		if !exists {
			c.UnAuthorized("未找到应用信息")
			return
		}

		appId := appIdVal.(string)

		// 获取请求信息
		method := c.Request.Method
		path := c.Request.URL.Path

		// 解析模块名和API路径
		parts := strings.Split(strings.Trim(path, "/"), "/")
		if len(parts) < 2 {
			c.Fail("无效的API路径")
			return
		}

		// 模块名（例如"user"）
		module := parts[1]

		// 检查权限
		err := svc.CheckAppPermission(appId, module, path, method)
		if err != nil {
			c.Fail(err.Error())
			return
		}

		// 权限验证通过，继续处理请求
		c.Next()
	}
}
