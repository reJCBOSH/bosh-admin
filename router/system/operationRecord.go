package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"
	"bosh-admin/module/system/operationRecord"

	"github.com/gin-gonic/gin"
)

func SetOperationRecordApiRouter(router *gin.RouterGroup) {
	group := router.Group("/operationRecord")
	groupRecord := router.Group("/operationRecord", middleware.OperationRecord())

	handler := operationRecord.NewSysOperationRecordApi()
	{
		group.GET("/list", ctx.Handler(handler.GetOperationRecordList))
		group.GET("/info", ctx.Handler(handler.GetOperationRecordInfo))
	}
	{
		groupRecord.POST("/del", ctx.Handler(handler.DelOperationRecord))
		groupRecord.POST("/batchDel", ctx.Handler(handler.BatchDelOperationRecord))
	}
}
