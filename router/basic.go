package router

import (
	"bosh-admin/core/ctx"
	"bosh-admin/module/basic"

	"github.com/gin-gonic/gin"
)

func SetBasicApiRouter(router *gin.RouterGroup) {
	group := router.Group("/basic")

	handler := basic.NewBasicApi()
	{
		group.GET("/captcha", ctx.Handler(handler.Captcha))
	}
}
