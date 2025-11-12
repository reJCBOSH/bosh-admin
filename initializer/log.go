package initializer

import (
	"os"

	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLog 初始化日志
func InitLog() {
	// 配置解码器
	encoder := log.CustomEncoder()
	// 日志按级别分开输出
	var cores []zapcore.Core
	startLevel := zapcore.DebugLevel
	if util.IsProd() {
		startLevel = zapcore.InfoLevel
	}
	for level := startLevel; level <= zapcore.ErrorLevel; level++ {
		writerSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(log.CustomLoggerWriter("_"+level.String())))
		var levelEnablerFunc zap.LevelEnablerFunc = func(l zapcore.Level) bool {
			if level == zapcore.ErrorLevel {
				return l >= level
			} else {
				return l == level
			}
		}
		core := zapcore.NewCore(encoder, writerSyncer, levelEnablerFunc)
		cores = append(cores, core)
	}
	// 日志配置
	options := []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1)}
	global.Logger = zap.New(zapcore.NewTee(cores...), options...).Sugar()
}
