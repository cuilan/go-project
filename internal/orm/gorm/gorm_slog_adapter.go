package gorm

import (
	"context"
	"log/slog"
	"time"

	"gorm.io/gorm/logger"
)

// slogLogger 是GORM与slog整合的日志适配器
type slogLogger struct {
	level logger.LogLevel
}

// LogMode 设置日志级别
func (l *slogLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.level = level
	return &newLogger
}

// Info 输出信息级别日志
func (l *slogLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Info {
		slog.Info("GORM", "msg", msg, "data", data)
	}
}

// Warn 输出警告级别日志
func (l *slogLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Warn {
		slog.Warn("GORM", "msg", msg, "data", data)
	}
}

// Error 输出错误级别日志
func (l *slogLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Error {
		slog.Error("GORM", "msg", msg, "data", data)
	}
}

// Trace 输出SQL跟踪日志
func (l *slogLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	// 构建日志字段
	fields := []interface{}{
		"elapsed", elapsed,
		"rows", rows,
	}

	if err != nil {
		fields = append(fields, "error", err)
		slog.Error("GORM SQL", fields...)
	} else {
		fields = append(fields, "sql", sql)
		slog.Info("GORM SQL", fields...)
	}
}
