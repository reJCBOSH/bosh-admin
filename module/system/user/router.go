package user

import (
	"bosh-admin/core/ctx"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/user")
	groupRecord := router.Group("/user", middleware.OperationRecord())

	user := NewSysUserApi()
	{
		group.GET("/list", ctx.Handler(user.GetUserList))
		group.GET("/info", ctx.Handler(user.GetUserInfo))
	}
	{
		groupRecord.POST("/add", ctx.Handler(user.AddUser))
		groupRecord.POST("/edit", ctx.Handler(user.EditUser))
		groupRecord.POST("/del", ctx.Handler(user.DelUser))
		groupRecord.POST("/resetPassword", ctx.Handler(user.ResetPassword))
		groupRecord.POST("/setStatus", ctx.Handler(user.SetUserStatus))
		groupRecord.POST("/editSelfInfo", ctx.Handler(user.EditSelfInfo))
		groupRecord.POST("/editSelfPassword", ctx.Handler(user.EditSelfPassword))
	}
}
