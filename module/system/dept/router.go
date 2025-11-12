package dept

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/dept")
	groupRecord := router.Group("/dept", middleware.OperationRecord())

	dept := NewSysDeptApi()
	{
		group.GET("/tree", ctx.Handler(dept.GetDeptTree))
		group.GET("/list", ctx.Handler(dept.GetDeptList))
		group.GET("/info", ctx.Handler(dept.GetDeptInfo))
	}
	{
		groupRecord.POST("/add", ctx.Handler(dept.AddDept))
		groupRecord.POST("/edit", ctx.Handler(dept.EditDept))
		groupRecord.POST("/del", ctx.Handler(dept.DelDept))
	}
}
