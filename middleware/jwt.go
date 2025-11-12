package middleware

import (
	"bosh-admin/core/ctx"
	"bosh-admin/core/log"
	"bosh-admin/model"
	"bosh-admin/module/auth"
	"bosh-admin/module/system/user"

	"github.com/gin-gonic/gin"
)

// JWTApiAuth JWT Api鉴权中间件
func JWTApiAuth() gin.HandlerFunc {
	return ctx.Handler(func(c *ctx.Context) {
		jwtSvc := auth.NewJWTSvc()
		// 获取access token
		accessToken, err := jwtSvc.GetAccessToken(c)
		if err != nil {
			c.UnAuthorized(err.Error())
			c.Abort()
			return
		}
		// 解析token
		claims, err := jwtSvc.ParseUserAccessToken(accessToken)
		if err != nil {
			c.UnAuthorized("无效的访问令牌")
			c.Abort()
			return
		}
		// 验证token
		err = jwtSvc.TokenValidate(claims.RegisteredClaims, auth.JwtSubjectAccess, auth.JwtAudienceApi)
		if err != nil {
			c.UnAuthorized("令牌验证失败")
			c.Abort()
			return
		}
		userSvc := user.NewSysUserSvc()
		var userInfo *model.SysUser
		userInfo, err = userSvc.GetUserById(claims.User.UserId)
		if err != nil {
			log.Error(err)
			c.UnAuthorized(err.Error())
			c.Abort()
			return
		}
		if userInfo.Status == 0 {
			c.UnAuthorized("用户已冻结")
			c.Abort()
			return
		}
		// 将claims存储到上下文中，供后续使用
		c.Set("userAccessClaims", claims)
		// 鉴权通过，继续处理请求
		c.Next()
	})
}
