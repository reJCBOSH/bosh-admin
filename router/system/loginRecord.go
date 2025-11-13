package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"
	"bosh-admin/module/system/loginRecord"

	"github.com/gin-gonic/gin"
)

func SetLoginRecordApiRouter(router *gin.RouterGroup) {
	group := router.Group("/loginRecord")
	groupRecord := router.Group("/loginRecord", middleware.OperationRecord())

	handler := loginRecord.NewSysLoginRecordApi()
	{
		group.GET("/list", ctx.Handler(handler.GetLoginRecordList))
	}
	{
		groupRecord.POST("/del", ctx.Handler(handler.DelLoginRecord))
		groupRecord.POST("/batchDel", ctx.Handler(handler.BatchDelLoginRecord))
	}
}
