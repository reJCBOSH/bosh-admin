package system

import (
	"github.com/gin-gonic/gin"
)

func SetSystemRouter(router *gin.RouterGroup) {
	group := router.Group("/system")

	SetApiRouter(group)
	SetDeptRouter(group)
	SetLoginRecordRouter(group)
	SetMenuRouter(group)
	SetOperationRecordRouter(group)
	SetRoleRouter(group)
	SetUserRouter(group)
}
