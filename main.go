package main

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/initializer"
	"bosh-admin/util"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化配置
	initializer.InitConfig()
	// 初始化日志
	initializer.InitLog()
	// 初始化数据库
	initializer.InitDB()
	// 程序结束前关闭数据库连接
	defer func() {
		if global.GormDB != nil {
			db, _ := global.GormDB.DB()
			_ = db.Close()
		}
	}()
	// 初始化IP2Region
	initializer.InitIP2Region()
	// 初始化参数校验
	initializer.InitValidator()
	// 初始化定时任务
	if util.IsProd() {
		initializer.InitCron()
	}
	// 初始化websocket
	initializer.InitWebsocket()
	// 初始化引擎
	initializer.InitEngine()

	addr := fmt.Sprintf(":%d", global.Config.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: global.Engine,
	}
	log.Info("服务已启动", addr)

	// 优雅关机
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("监听: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Info("关闭服务中...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("服务关闭失败: ", err)
	}

	log.Info("服务已退出")
}
