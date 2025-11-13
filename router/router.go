package router

import (
	"net/http"

	"bosh-admin/core/ctx"
	"bosh-admin/global"
	"bosh-admin/middleware"
	"bosh-admin/router/system"

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
		SetAuthApiRouter(public)
		SetBasicApiRouter(public)
	}

	private := group.Group("", middleware.JWTApiAuth())
	{
		system.SetSystemApiRouter(private)
	}
}

func SetOpenapiRouter(engine *gin.Engine) {

}
