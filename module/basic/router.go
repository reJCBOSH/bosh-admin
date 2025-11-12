package basic

import (
	"bosh-admin/core/ctx"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/basic")

	basic := NewBasicApi()
	{
		group.GET("/captcha", ctx.Handler(basic.Captcha))
	}
}
