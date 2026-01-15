package initializer

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/service/sse"
)

// InitSSE 初始化SSE
func InitSSE() {
	global.SSESvc = sse.NewSSEService()
	log.Info("SSE初始化完成")
}
