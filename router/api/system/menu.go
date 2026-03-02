package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/domain/api/system/menu"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetMenuRouter(router *gin.RouterGroup) {
	group := router.Group("/menu")
	groupRecord := router.Group("/menu", middleware.OperationRecord())

	handler := menu.NewHandlerSysMenu()
	{
		group.GET("/tree", ctx.Handler(handler.GetMenuTree))
		group.GET("/list", ctx.Handler(handler.GetMenuList))
		group.GET("/info", ctx.Handler(handler.GetMenuInfo))
		group.GET("/asyncRoutes", ctx.Handler(handler.GetAsyncRoutes))
	}
	{
		groupRecord.POST("/add", ctx.Handler(handler.AddMenu))
		groupRecord.POST("/edit", ctx.Handler(handler.EditMenu))
		groupRecord.POST("/del", ctx.Handler(handler.DelMenu))
	}
}
