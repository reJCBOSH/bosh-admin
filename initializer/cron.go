package initializer

import (
	"bosh-admin/core/log"

	"github.com/robfig/cron/v3"
)

// InitCron 初始化定时任务
func InitCron() {
	var err error
	c := cron.New(cron.WithSeconds())
	// 每日0点执行
	_, err = c.AddFunc("0 0 0 * * ?", func() {
		// 定时任务逻辑
	})
	if err != nil {
		log.Error("每日0点执行定时任务失败", err)
	}
	// 启动定时任务
	c.Start()
}
