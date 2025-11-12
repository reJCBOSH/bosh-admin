package auth

import (
	"bosh-admin/core/ctx"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/auth")

	auth := NewAuthApi()
	{
		group.POST("/user/login", ctx.Handler(auth.UserLogin))
		group.POST("/refreshToken", ctx.Handler(auth.RefreshToken))
	}
}
