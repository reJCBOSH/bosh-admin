package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/domain/api/auth"

	"github.com/gin-gonic/gin"
)

func SetAuthRouter(router *gin.RouterGroup) {
	group := router.Group("/auth")

	handler := auth.NewHandlerAuth()
	{
		group.POST("/user/login", ctx.Handler(handler.UserLogin))
		group.POST("/refreshToken", ctx.Handler(handler.RefreshToken))
	}
}
