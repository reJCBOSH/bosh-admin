package initializer

import (
	"fmt"
	"os"

	"bosh-admin/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	// 设置配置文件路径
	configFile := "config.yaml"
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		configFile = configEnv
	}
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件错误: %s \n", err.Error()))
	}
	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件更新: ", e.Name)
		// 重载配置
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println("重载配置失败:", err.Error())
		}
	})
	// 将配置赋值给配置变量
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println("更新配置失败:", err.Error())
	}
}
