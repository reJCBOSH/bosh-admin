package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"
	"bosh-admin/module/system/role"

	"github.com/gin-gonic/gin"
)

func SetRoleApiRouter(router *gin.RouterGroup) {
	group := router.Group("/role")
	groupRecord := router.Group("/role", middleware.OperationRecord())

	handler := role.NewSysRoleApi()
	{
		group.GET("/list", ctx.Handler(handler.GetRoleList))
		group.GET("/info", ctx.Handler(handler.GetRoleInfo))
		group.GET("/menu", ctx.Handler(handler.GetRoleMenu))
		group.GET("/menuIds", ctx.Handler(handler.GetRoleMenuIds))
		group.GET("/deptIds", ctx.Handler(handler.GetRoleDeptIds))
	}
	{
		groupRecord.POST("/add", ctx.Handler(handler.AddRole))
		groupRecord.POST("/edit", ctx.Handler(handler.EditRole))
		groupRecord.POST("/del", ctx.Handler(handler.DelRole))
		groupRecord.POST("/setMenuAuth", ctx.Handler(handler.SetRoleMenuAuth))
		groupRecord.POST("/setDataAuth", ctx.Handler(handler.SetRoleDataAuth))
		groupRecord.POST("/setStatus", ctx.Handler(handler.SetRoleStatus))
	}
}
