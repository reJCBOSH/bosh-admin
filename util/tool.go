package util

import (
	"bosh-admin/core/log"
	"bosh-admin/global"
	"strings"
)

// IsProd 是否生产环境
func IsProd() bool {
	if global.Config.Server.Env == global.PROD {
		return true
	}
	return false
}

func IsDev() bool {
	if global.Config.Server.Env == global.DEV {
		return true
	}
	return false
}

// IP2Region IP转地理位置
func IP2Region(ip string) string {
	result, err := global.XdbSearcher.SearchByStr(ip)
	if err != nil {
		log.Error("查询IP地理信息失败:", err)
		return ""
	}
	resultArr := strings.Split(result, "|")
	// 返回省份城市
	if resultArr[2] != "0" && resultArr[3] != "0" {
		return resultArr[2] + resultArr[3]
	}
	// 返回国家省份
	if resultArr[0] != "0" && resultArr[2] != "0" && resultArr[3] == "0" {
		return resultArr[0] + resultArr[3]
	}
	// 返回国家
	if resultArr[0] != "0" && resultArr[2] == "0" && resultArr[3] == "0" {
		return resultArr[0]
	}
	// 内网IP场景
	if resultArr[0] == "0" && resultArr[2] == "0" && resultArr[3] != "0" {
		return resultArr[3]
	}
	return ""
}
