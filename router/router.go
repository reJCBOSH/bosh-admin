package router

import (
	"bosh-admin/middleware"
	"bosh-admin/module/auth"
	"bosh-admin/module/basic"
	"net/http"

	"bosh-admin/core/ctx"
	"bosh-admin/global"
	"bosh-admin/module/system"

	"github.com/gin-gonic/gin"
)

func SetHealthRouter(engine *gin.Engine) {
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": ctx.SUCCESS,
		})
	})
}

func SetStaticRouter(engine *gin.Engine) {
	if global.Config.Local.Path != "" && global.Config.Local.StorePath != "" {
		engine.Static(global.Config.Local.Path, global.Config.Local.StorePath)
	} else {
		engine.Static("/static", "static")
	}
}

func SetWebSocketRouter(engine *gin.Engine) {
	engine.GET("/ws", ctx.Handler(func(c *ctx.Context) {
		global.WsHub.HandleConnection(c.Writer, c.Request, c.Query("token"))
	}))
}

func SetApiRouter(engine *gin.Engine) {
	group := engine.Group("/api")

	public := group.Group("")
	{
		auth.SetApiRouter(public)
		basic.SetApiRouter(public)
	}

	private := group.Group("", middleware.JWTApiAuth())
	{
		system.SetApiRouter(private)
	}
}

func SetOpenapiRouter(engine *gin.Engine) {

}
