package loginRecord

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/loginRecord")
	groupRecord := router.Group("/loginRecord", middleware.OperationRecord())

	loginRecord := NewSysLoginRecordApi()
	{
		group.GET("/list", ctx.Handler(loginRecord.GetLoginRecordList))
	}
	{
		groupRecord.POST("/del", ctx.Handler(loginRecord.DelLoginRecord))
		groupRecord.POST("/batchDel", ctx.Handler(loginRecord.BatchDelLoginRecord))
	}
}
