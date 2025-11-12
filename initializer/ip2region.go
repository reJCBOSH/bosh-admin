package initializer

import (
	"fmt"

	"bosh-admin/global"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

// InitIP2Region 初始化ip2region
func InitIP2Region() {
	dbPath := "ip2region_v4.xdb"
	version := xdb.IPv4
	err := xdb.VerifyFromFile(dbPath)
	if err != nil {
		panic(fmt.Sprintf("验证ip2region.xdb文件失败: %v", err))
	}
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		panic(fmt.Sprintf("加载ip2region.xdb文件失败: %v", err))
	}
	global.XdbSearcher, err = xdb.NewWithBuffer(version, cBuff)
	if err != nil {
		panic(fmt.Sprintf("初始化XdbSearcher失败: %v", err))
	}
}
