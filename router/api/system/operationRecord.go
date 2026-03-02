package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/domain/api/system/operationRecord"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetOperationRecordRouter(router *gin.RouterGroup) {
	group := router.Group("/operationRecord")
	groupRecord := router.Group("/operationRecord", middleware.OperationRecord())

	handler := operationRecord.NewHandlerSysOperationRecord()
	{
		group.GET("/list", ctx.Handler(handler.GetOperationRecordList))
		group.GET("/info", ctx.Handler(handler.GetOperationRecordInfo))
	}
	{
		groupRecord.POST("/del", ctx.Handler(handler.DelOperationRecord))
		groupRecord.POST("/batchDel", ctx.Handler(handler.BatchDelOperationRecord))
	}
}
