package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorReset  = "\033[0m"
)

// CustomTextHandler 实现了 slog.Handler 接口，用于输出不带 key 的文本日志格式。
type CustomTextHandler struct {
	opts     slog.HandlerOptions
	writer   io.Writer
	attrs    []slog.Attr
	groups   []string
	colorize bool // 是否为日志级别启用颜色
}

// NewCustomTextHandler 创建一个新的 CustomTextHandler。
func NewCustomTextHandler(w io.Writer, opts *slog.HandlerOptions, colorize bool) *CustomTextHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &CustomTextHandler{
		opts:     *opts,
		writer:   w,
		colorize: colorize,
	}
}

// Enabled 检查该日志级别是否启用。
func (h *CustomTextHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return level >= minLevel
}

// Handle 处理并格式化日志记录。
func (h *CustomTextHandler) Handle(_ context.Context, r slog.Record) error {
	var buf bytes.Buffer

	// 格式化时间
	if !r.Time.IsZero() {
		fmt.Fprintf(&buf, "[%s] ", r.Time.Format("2006-01-02 15:04:05.000"))
	}

	// 格式化级别
	var levelStr string
	var color string
	switch r.Level {
	case slog.LevelDebug:
		levelStr = "[D]"
		color = colorGreen
	case slog.LevelInfo:
		levelStr = "[I]"
		color = colorBlue
	case slog.LevelWarn:
		levelStr = "[W]"
		color = colorYellow
	case slog.LevelError:
		levelStr = "[E]"
		color = colorRed
	default:
		levelStr = fmt.Sprintf("[%s]", r.Level.String())
	}

	if h.colorize && color != "" {
		fmt.Fprintf(&buf, "%s%s%s ", color, levelStr, colorReset)
	} else {
		fmt.Fprintf(&buf, "%s ", levelStr)
	}

	// 格式化源码位置
	if h.opts.AddSource && r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		// 仅保留文件名
		_, fileName := path.Split(f.File)
		fmt.Fprintf(&buf, "[%s:%d] ", fileName, f.Line)
	}

	// 写入日志消息
	buf.WriteString(r.Message)

	// --- 处理并格式化所有属性 ---
	var attrs []slog.Attr
	// 添加由 WithAttrs 创建的属性
	attrs = append(attrs, h.attrs...)
	// 添加日志记录本身的属性
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	// 将分组和属性格式化为 key=value
	if len(attrs) > 0 {
		buf.WriteByte(' ')
		for i, attr := range attrs {
			h.formatAttr(&buf, attr, h.groups)
			if i < len(attrs)-1 {
				buf.WriteByte(' ')
			}
		}
	}

	buf.WriteByte('\n')
	_, err := h.writer.Write(buf.Bytes())
	return err
}

// formatAttr 递归地将属性（包括分组内的）格式化为 key=value。
func (h *CustomTextHandler) formatAttr(buf *bytes.Buffer, a slog.Attr, groups []string) {
	// 如果是分组，则递归处理
	if a.Value.Kind() == slog.KindGroup {
		newGroups := append(groups, a.Key)
		for i, gAttr := range a.Value.Group() {
			h.formatAttr(buf, gAttr, newGroups)
			if i < len(a.Value.Group())-1 {
				buf.WriteByte(' ')
			}
		}
		return
	}

	// 格式化普通属性
	key := a.Key
	if len(groups) > 0 {
		key = strings.Join(groups, ".") + "." + key
	}
	fmt.Fprintf(buf, "%s=%s", key, a.Value)
}

// WithAttrs 返回一个新的 Handler，它包含已有的和新增的属性。
func (h *CustomTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := *h
	newHandler.attrs = append(h.attrs, attrs...)
	return &newHandler
}

// WithGroup 返回一个新的 Handler，它包含新增的分组。
func (h *CustomTextHandler) WithGroup(name string) slog.Handler {
	newHandler := *h
	newHandler.groups = append(h.groups, name)
	return &newHandler
}

// MultiHandler 是一个 slog.Handler，它可以将日志记录分发到多个子 Handler。
// 每个子 Handler 根据自己的 Enabled 方法决定是否处理该日志。
type MultiHandler struct {
	handlers []slog.Handler
}

// NewMultiHandler 创建一个新的 MultiHandler。
func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

// Enabled 检查是否有任何一个子 Handler 对指定的日志级别启用了。
func (h *MultiHandler) Enabled(_ context.Context, level slog.Level) bool {
	// 实际上，这个 Enabled 方法对于组合处理器来说意义不大，
	// 因为决策权在子 Handler 手中。我们让它默认返回 true，
	// 然后在 Handle 方法中进行真正的判断。
	return true
}

// Handle 将日志记录分发给所有对其级别启用了的子 Handler。
func (h *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, r.Level) {
			if err := handler.Handle(ctx, r); err != nil {
				// 理论上，一个 handler 失败不应阻止其他 handler 工作，
				// 但为了简单起见，我们返回第一个遇到的错误。
				// 实际应用中可以考虑更复杂的错误处理策略。
				return err
			}
		}
	}
	return nil
}

// WithAttrs 为所有子 Handler 添加属性。
func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithAttrs(attrs)
	}
	return NewMultiHandler(newHandlers...)
}

// WithGroup 为所有子 Handler 添加分组。
func (h *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithGroup(name)
	}
	return NewMultiHandler(newHandlers...)
}

// parseLevel 将字符串级别转换为 slog.Level，并提供默认值。
func parseLevel(levelStr string, defaultLevel slog.Level) slog.Level {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(levelStr)); err != nil {
		return defaultLevel
	}
	return lvl
}

// createHandler 根据输出器、格式和级别等配置创建一个 slog.Handler。
func createHandler(writer io.Writer, format string, level slog.Level, addSource bool) slog.Handler {
	opts := &slog.HandlerOptions{
		AddSource: addSource,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05.000"))
				}
			}
			if a.Key == slog.SourceKey {
				if source, ok := a.Value.Any().(*slog.Source); ok {
					source.File = path.Base(source.File)
				}
			}
			return a
		},
	}

	switch strings.ToLower(format) {
	case "json":
		return slog.NewJSONHandler(writer, opts)
	default: // 默认为 text 格式
		// 仅当输出目标是控制台时才启用颜色
		isConsole := writer == os.Stdout
		return NewCustomTextHandler(writer, opts, isConsole)
	}
}

// InitLogger 根据传入的详细配置初始化全局日志记录器。
func InitLogger(cfg *LoggerConfig) {
	var handlers []slog.Handler

	// --- 配置控制台输出 ---
	if cfg.Console.EnableConsole {
		level := parseLevel(cfg.Console.Level, slog.LevelInfo)
		handler := createHandler(os.Stdout, cfg.Console.ConsoleFormat, level, cfg.Console.AddSource)
		handlers = append(handlers, handler)
	}

	// --- 配置-文件输出 ---
	if cfg.File.EnableFile {
		// 设置 lumberjack 的默认值
		if cfg.File.Path == "" {
			cfg.File.Path = "./logs"
		}
		if cfg.File.Filename == "" {
			cfg.File.Filename = "default.log"
		}
		if cfg.File.FileMaxSize == 0 {
			cfg.File.FileMaxSize = 100
		}
		if cfg.File.FileMaxAge == 0 {
			cfg.File.FileMaxAge = 3
		}
		if cfg.File.FileMaxBackups == 0 {
			cfg.File.FileMaxBackups = 3
		}

		fileWriter := &lumberjack.Logger{
			Filename:   path.Join(cfg.File.Path, cfg.File.Filename),
			MaxSize:    cfg.File.FileMaxSize,
			MaxAge:     cfg.File.FileMaxAge,
			MaxBackups: cfg.File.FileMaxBackups,
			LocalTime:  true,
			Compress:   cfg.File.FileCompress,
		}
		level := parseLevel(cfg.File.Level, slog.LevelInfo)
		handler := createHandler(fileWriter, cfg.File.FileFormat, level, cfg.File.AddSource)
		handlers = append(handlers, handler)
	}

	if len(handlers) == 0 {
		slog.Warn("logger is not configured for any output")
		return
	}

	// 使用 MultiHandler 将多个处理器组合起来
	combinedHandler := NewMultiHandler(handlers...)
	slog.SetDefault(slog.New(combinedHandler))

	slog.Info("logger initialization completed", "handlers", len(handlers))
	if cfg.Console.Level == "debug" {
		consoleJsonCfg, _ := json.Marshal(cfg.Console)
		slog.Debug("logger config", "console", string(consoleJsonCfg))
		fileJsonCfg, _ := json.Marshal(cfg.File)
		slog.Debug("logger config", "file", string(fileJsonCfg))
	}
}
