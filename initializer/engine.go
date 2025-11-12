package initializer

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/middleware"
	"bosh-admin/router"
	"bosh-admin/util"

	"github.com/gin-gonic/gin"
)

// InitEngine 初始化引擎
func InitEngine() {
	if util.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	// 跨域中间件
	engine.Use(middleware.Cors())
	// 使用gin默认Logger、Recovery中间件
	engine.Use(gin.Logger(), gin.Recovery())

	router.SetHealthRouter(engine)
	router.SetStaticRouter(engine)
	router.SetWebSocketRouter(engine)
	router.SetApiRouter(engine)
	log.Info("路由注册完成")

	global.Engine = engine
}
