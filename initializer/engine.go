package initializer

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/middleware"
	"bosh-admin/router"
	"bosh-admin/util"

	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	// 安全页眉中间件
	engine.Use(ginhelmet.Default())
	engine.Use(middleware.Prometheus())

	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.SetHealthRouter(engine)
	router.SetStaticRouter(engine)
	router.SetWebSocketRouter(engine)
	router.SetSSERouter(engine)
	router.SetApiRouter(engine)
	log.Info("路由注册完成")

	global.Engine = engine
}
