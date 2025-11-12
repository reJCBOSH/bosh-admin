package menu

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/menu")
	groupRecord := router.Group("/menu", middleware.OperationRecord())

	menu := NewSysMenuApi()
	{
		group.GET("/tree", ctx.Handler(menu.GetMenuTree))
		group.GET("/list", ctx.Handler(menu.GetMenuList))
		group.GET("/info", ctx.Handler(menu.GetMenuInfo))
		group.GET("/asyncRoutes", ctx.Handler(menu.GetAsyncRoutes))
	}
	{
		groupRecord.POST("/add", ctx.Handler(menu.AddMenu))
		groupRecord.POST("/edit", ctx.Handler(menu.EditMenu))
		groupRecord.POST("/del", ctx.Handler(menu.DelMenu))
	}
}
