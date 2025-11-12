package role

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/role")
	groupRecord := router.Group("/role", middleware.OperationRecord())

	role := NewSysRoleApi()
	{
		group.GET("/list", ctx.Handler(role.GetRoleList))
		group.GET("/info", ctx.Handler(role.GetRoleInfo))
		group.GET("/menu", ctx.Handler(role.GetRoleMenu))
		group.GET("/menuIds", ctx.Handler(role.GetRoleMenuIds))
		group.GET("/deptIds", ctx.Handler(role.GetRoleDeptIds))
	}
	{
		groupRecord.POST("/add", ctx.Handler(role.AddRole))
		groupRecord.POST("/edit", ctx.Handler(role.EditRole))
		groupRecord.POST("/del", ctx.Handler(role.DelRole))
		groupRecord.POST("/setMenuAuth", ctx.Handler(role.SetRoleMenuAuth))
		groupRecord.POST("/setDataAuth", ctx.Handler(role.SetRoleDataAuth))
		groupRecord.POST("/setStatus", ctx.Handler(role.SetRoleStatus))
	}
}
