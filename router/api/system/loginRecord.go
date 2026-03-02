package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/domain/api/system/loginRecord"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetLoginRecordRouter(router *gin.RouterGroup) {
	group := router.Group("/loginRecord")
	groupRecord := router.Group("/loginRecord", middleware.OperationRecord())

	handler := loginRecord.NewHandlerSysLoginRecord()
	{
		group.GET("/list", ctx.Handler(handler.GetLoginRecordList))
	}
	{
		groupRecord.POST("/del", ctx.Handler(handler.DelLoginRecord))
		groupRecord.POST("/batchDel", ctx.Handler(handler.BatchDelLoginRecord))
	}
}
