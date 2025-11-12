package initializer

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/websocket"
)

// InitWebsocket 初始化websocket
func InitWebsocket() {
	global.WsHub = websocket.NewHub(global.Logger)
	go global.WsHub.Start()
	log.Info("websocket初始化完成")
}
