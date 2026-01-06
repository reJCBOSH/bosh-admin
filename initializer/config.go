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
	
	// 设置环境变量支持
	v.AutomaticEnv()
	
	// 从配置文件读取初始值
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件错误: %s \n", err.Error()))
	}
	
	// 优先使用环境变量覆盖配置文件中的值
	overrideConfigWithEnv(v)
	
	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件更新: ", e.Name)
		// 重载配置
		overrideConfigWithEnv(v)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println("重载配置失败:", err.Error())
		}
	})
	// 将配置赋值给配置变量
	overrideConfigWithEnv(v)
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println("更新配置失败:", err.Error())
	}
}

// overrideConfigWithEnv 使用环境变量覆盖配置值
func overrideConfigWithEnv(v *viper.Viper) {
	// 数据库配置环境变量覆盖
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		v.Set("mysql.ip", dbHost)
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		v.Set("mysql.port", dbPort)
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		v.Set("mysql.username", dbUser)
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		v.Set("mysql.password", dbPassword)
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		v.Set("mysql.database", dbName)
	}
}