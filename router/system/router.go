package system

import (
	"github.com/gin-gonic/gin"
)

func SetSystemApiRouter(router *gin.RouterGroup) {
	group := router.Group("/system")

	SetDeptApiRouter(group)
	SetLoginRecordApiRouter(group)
	SetMenuApiRouter(group)
	SetOperationRecordApiRouter(group)
	SetRoleApiRouter(group)
	SetUserApiRouter(group)
}
