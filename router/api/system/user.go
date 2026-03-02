package system

import (
	"bosh-admin/core/ctx"
	"bosh-admin/domain/api/system/user"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetUserRouter(router *gin.RouterGroup) {
	group := router.Group("/user")
	groupRecord := router.Group("/user", middleware.OperationRecord())

	handler := user.NewHandlerSysUser()
	{
		group.GET("/list", ctx.Handler(handler.GetUserList))
		group.GET("/info", ctx.Handler(handler.GetUserInfo))
		group.GET("/getSelfInfo", ctx.Handler(handler.GetSelfInfo))
	}
	{
		groupRecord.POST("/add", ctx.Handler(handler.AddUser))
		groupRecord.POST("/edit", ctx.Handler(handler.EditUser))
		groupRecord.POST("/del", ctx.Handler(handler.DelUser))
		groupRecord.POST("/resetPassword", ctx.Handler(handler.ResetPassword))
		groupRecord.POST("/setStatus", ctx.Handler(handler.SetUserStatus))
		groupRecord.POST("/editSelfInfo", ctx.Handler(handler.EditSelfInfo))
		groupRecord.POST("/editSelfPassword", ctx.Handler(handler.EditSelfPassword))
	}
}
