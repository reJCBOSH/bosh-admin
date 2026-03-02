package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/domain/api/basic"

	"github.com/gin-gonic/gin"
)

func SetBasicRouter(router *gin.RouterGroup) {
	group := router.Group("/basic")

	handler := basic.NewHandlerBasic()
	{
		group.GET("/captcha", ctx.Handler(handler.Captcha))
		group.GET("/publicKey", ctx.Handler(handler.GetPublicKey))
	}
}
