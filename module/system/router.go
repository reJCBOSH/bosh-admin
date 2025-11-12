package system

import (
	"bosh-admin/module/system/dept"
	"bosh-admin/module/system/loginRecord"
	"bosh-admin/module/system/menu"
	"bosh-admin/module/system/operationRecord"
	"bosh-admin/module/system/role"
	"bosh-admin/module/system/user"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.RouterGroup) {
	group := router.Group("/system")

	dept.SetApiRouter(group)
	loginRecord.SetApiRouter(group)
	menu.SetApiRouter(group)
	operationRecord.SetApiRouter(group)
	role.SetApiRouter(group)
	user.SetApiRouter(group)
}
