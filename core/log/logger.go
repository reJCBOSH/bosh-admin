package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"bosh-admin/global"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
)

// customEncodeTime 自定义时间编码器
func customEncodeTime() zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(global.Config.Log.TimestampFormat))
	}
}

// getProjectRoot 获取项目根目录
func getProjectRoot() string {
	// 获取调用此函数的文件的路径（通常为日志配置文件）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	dir := filepath.Dir(filename)
	for {
		// 检查当前目录是否包含go.mod
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			fmt.Println(dir)
			return dir
		}
		// 向上级目录查找
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break // 到达文件系统根目录
		}
		dir = parentDir
	}
	return "" // 未找到
}

// customEncodeCaller 自定义调用解码器
func customEncodeCaller(basePath string) zapcore.CallerEncoder {
	return func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// 获取相对路径
		relPath, err := filepath.Rel(basePath, caller.File)
		if err != nil {
			relPath = caller.File // 出错则使用原路径
		} else {
			relPath = strings.Replace(relPath, "\\", "/", -1)
		}
		// 格式化为"文件:行号"
		enc.AppendString(fmt.Sprintf("%s:%d", relPath, caller.Line))
	}
}

// CustomEncoder 自定义解码器
func CustomEncoder() zapcore.Encoder {
	var encoder zapcore.Encoder
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "M",
		LevelKey:       "L",
		TimeKey:        "T",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    "F",
		StacktraceKey:  zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customEncodeTime(),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   customEncodeCaller(getProjectRoot()),
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 解码器格式
	if strings.ToLower(global.Config.Log.Format) == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

// CustomLogger 自定义日志
type CustomLogger struct {
	logger *lumberjack.Logger
	date   string // 日志日期
	suffix string // 日志后缀
}

// Write 自定义日志Write
func (l *CustomLogger) Write(p []byte) (n int, err error) {
	now := time.Now().Local()
	nowDate := now.Format(time.DateOnly)
	if l.date != nowDate {
		defer func(logger *lumberjack.Logger) {
			_ = logger.Close()
		}(l.logger)
		l.date = nowDate
		l.logger.Filename = fmt.Sprintf("%s/%s/%s/%s%s.log", global.Config.Log.RootDir, now.Format("2006-01"), nowDate, nowDate, l.suffix)
	}
	return l.logger.Write(p)
}

// CustomLoggerWriter 自定义日志写入器
func CustomLoggerWriter(suffix string) *CustomLogger {
	now := time.Now().Local()
	nowDate := now.Format(time.DateOnly)
	return &CustomLogger{
		logger: &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s/%s/%s%s.log", global.Config.Log.RootDir, now.Format("2006-01"), nowDate, nowDate, suffix),
			MaxSize:    global.Config.Log.MaxSize,    // 每个日志文件保存的最大尺寸 单位：MB
			MaxAge:     global.Config.Log.MaxAge,     // 文件最多保存多少天
			MaxBackups: global.Config.Log.MaxBackups, // 日志文件最多保存多少个备份
			Compress:   global.Config.Log.Compress,   // 是否压缩
		},
		date:   nowDate,
		suffix: suffix,
	}
}

// CustomGormLogger 自定义Gorm日志
func CustomGormLogger(suffix string) logger.Interface {
	return logger.New(log.New(CustomLoggerWriter(suffix), "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
		LogLevel:                  logger.Info,
	})
}
