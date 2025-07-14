package logger

import (
	"io"
	"log/slog"
	"os"
	"path"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger 初始化 slog.Logger
// 配置 logger 将日志同时输出到控制台和文件（如果启用）
// 日志轮转功能由 lumberjack 库处理
func InitLogger(logPath, appName, level string, enableToFile bool, maxAge int) {
	var writers []io.Writer

	// 默认添加控制台输出
	writers = append(writers, os.Stdout)

	// 如果启用文件写入，则添加文件写入器，并配置日志轮转
	if enableToFile {
		writers = append(writers, &lumberjack.Logger{
			Filename:   path.Join(logPath, appName+".log"), // 日志文件路径
			MaxSize:    100,                                // 单个文件最大体积（单位：MB）
			MaxAge:     maxAge,                             // 旧日志文件保留天数
			MaxBackups: 3,                                  // 最多保留的旧日志文件数量
			LocalTime:  true,                               // 使用本地时间作为时间戳
			Compress:   false,                              // 不压缩旧的日志文件
		})
	}

	// 创建一个多路写入器，将日志同时写入所有指定的 writers
	multiWriter := io.MultiWriter(writers...)

	// 从字符串解析日志级别
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		// 如果解析失败，默认为 Info 级别
		lvl = slog.LevelInfo
	}

	// 创建一个新的 JSON 处理器，它会将日志记录格式化为 JSON
	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		// 在日志输出中添加源码位置（文件名和行号）
		AddSource: true,
		// 设置日志记录的最低级别
		Level: lvl,
		// ReplaceAttr 自定义或修改日志属性
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 格式化时间
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05.000"))
				}
			}
			// 精简文件路径
			if a.Key == slog.SourceKey {
				if source, ok := a.Value.Any().(*slog.Source); ok {
					source.File = path.Base(source.File)
					source.Function = path.Base(source.Function)
				}
			}
			return a
		},
	})

	// 将新创建的 logger 设置为全局默认 logger
	slog.SetDefault(slog.New(handler))

	// 记录一条日志，表示 logger 初始化完成
	slog.Info("logger initialization completed", "level", level, "enableToFile", enableToFile)
}
