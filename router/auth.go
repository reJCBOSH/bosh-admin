package router

import (
	"bosh-admin/core/ctx"
	"bosh-admin/module/auth"

	"github.com/gin-gonic/gin"
)

func SetAuthApiRouter(router *gin.RouterGroup) {
	group := router.Group("/auth")

	handler := auth.NewAuthApi()
	{
		group.POST("/user/login", ctx.Handler(handler.UserLogin))
		group.POST("/refreshToken", ctx.Handler(handler.RefreshToken))
	}
}
