package operationRecord

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/operationRecord")
	groupRecord := router.Group("/operationRecord", middleware.OperationRecord())

	operationRecord := NewSysOperationRecordApi()
	{
		group.GET("/list", ctx.Handler(operationRecord.GetOperationRecordList))
		group.GET("/info", ctx.Handler(operationRecord.GetOperationRecordInfo))
	}
	{
		groupRecord.POST("/del", ctx.Handler(operationRecord.DelOperationRecord))
		groupRecord.POST("/batchDel", ctx.Handler(operationRecord.BatchDelOperationRecord))
	}
}
