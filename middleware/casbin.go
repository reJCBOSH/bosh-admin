package middleware

import (
	"bosh-admin/core/ctx"
	"bosh-admin/global"
	"bosh-admin/module/auth"

	"github.com/gin-gonic/gin"
)

// CasbinRBAC 接口访问权限中间件
func CasbinRBAC() gin.HandlerFunc {
	return ctx.Handler(func(c *ctx.Context) {
		if global.Config.Server.Env == global.DEV {
			c.Next()
		} else {
			jwtSvc := auth.NewJWTSvc()
			waitUser := jwtSvc.GetUserClaims(c)
			if waitUser.RoleCode != global.SuperAdmin {
				// 获取用户当前角色
				sub := waitUser.RoleId
				// 获取请求路径
				obj := c.Request.URL.Path
				// 获取请求方法
				act := c.Request.Method
				e := auth.NewCasbinSvc().Casbin()
				// 判断策略中是否存在
				success, _ := e.Enforce(sub, obj, act)
				if !success {
					c.Fail("访问权限不足")
					c.Abort()
					return
				}
			} else {
				c.Next()
			}
		}
	})
}
