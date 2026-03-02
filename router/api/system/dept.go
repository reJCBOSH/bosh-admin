package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/domain/api/system/dept"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetDeptRouter(router *gin.RouterGroup) {
	group := router.Group("/dept")
	groupRecord := router.Group("/dept", middleware.OperationRecord())

	handler := dept.NewHandlerSysDept()
	{
		group.GET("/tree", ctx.Handler(handler.GetDeptTree))
		group.GET("/list", ctx.Handler(handler.GetDeptList))
		group.GET("/info", ctx.Handler(handler.GetDeptInfo))
	}
	{
		groupRecord.POST("/add", ctx.Handler(handler.AddDept))
		groupRecord.POST("/edit", ctx.Handler(handler.EditDept))
		groupRecord.POST("/del", ctx.Handler(handler.DelDept))
	}
}
