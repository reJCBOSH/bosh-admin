package openapi

//func InitOpenAPIRouter() {
//	openAPI := NewOpenAPI()
//
//	// 外部接口路由组（需要API密钥验证）
//	openAPIRouter := router.Engine.Group("/openapi")
//	openAPIRouter.Use(middleware.OpenAPIAuth())
//	{
//		// 用户模块路由组（需要权限验证）
//		userRouter := openAPIRouter.Group("/user")
//		userRouter.Use(middleware.OpenAPIPermission())
//		{
//			userRouter.GET("/info", openAPI.GetUserInfo)
//			userRouter.GET("/list", openAPI.GetUserList)
//		}
//
//		// 示例接口 - 获取当前时间（无需特殊权限）
//		openAPIRouter.GET("/time", openAPI.GetCurrentTime)
//	}
//
//	// 内部管理路由组（需要JWT验证）
//	adminRouter := router.Engine.Group("/admin/openapi")
//	adminRouter.Use(middleware.JWTAuth())
//	{
//		// 应用管理
//		adminRouter.POST("/app", openAPI.CreateApp)
//
//		// API密钥管理
//		adminRouter.POST("/app-key", openAPI.CreateAppKey)
//
//		// 应用权限管理
//		adminRouter.POST("/app-perm", openAPI.CreateAppPerm)
//		adminRouter.PUT("/app-perm", openAPI.UpdateAppPerm)
//		adminRouter.DELETE("/app-perm/:id", openAPI.DeleteAppPerm)
//		adminRouter.GET("/app-perm/list", openAPI.GetAppPermList)
//	}
//}
