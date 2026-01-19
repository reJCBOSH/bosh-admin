package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"
	"bosh-admin/module/system/api"

	"github.com/gin-gonic/gin"
)

func SetApiApiRouter(router *gin.RouterGroup) {
	group := router.Group("/api")
	groupRecord := router.Group("/api", middleware.OperationRecord())

	handler := api.NewSysApiApi()
	{
		group.GET("/list", ctx.Handler(handler.GetApiList))
		group.GET("/info", ctx.Handler(handler.GetApiInfo))
		group.GET("/groups", ctx.Handler(handler.GetApiGroups))
	}
	{
		groupRecord.POST("/add", ctx.Handler(handler.AddApi))
		groupRecord.POST("/edit", ctx.Handler(handler.EditApi))
		groupRecord.POST("/del", ctx.Handler(handler.DelApi))
	}
}
