package initializer

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/sse"
)

// InitSSE 初始化SSE
func InitSSE() {
	global.SSESrv = sse.NewSSEService()
	log.Info("SSE初始化完成")
}
