package module

import (
	"log/slog"

	"github.com/spf13/viper"
)

// Module 是一个通用模块接口
type Module interface {
	// Name 返回模块名称
	Name() string
	// Init 根据 Viper 的配置初始化模块
	Init(v *viper.Viper) error
	// Close 释放模块资源
	Close() error
}

var modules []Module

// Register 注册一个或多个模块
func Register(m ...Module) {
	slog.Info("registering modules...", "count", len(m))
	for _, module := range m {
		modules = append(modules, module)
		slog.Debug("registered", "module", module.Name())
	}
	slog.Info("all modules registered")
}

// InitModules 初始化所有已注册的模块
func InitModules(v *viper.Viper) {
	slog.Info("initializing pluggable modules...", "count", len(modules))
	for _, m := range modules {
		// 跳过已被作为核心服务提前初始化的 logger 模块
		if m.Name() == "log" {
			continue
		}
		// 检查配置中是否存在该模块的配置项
		if v.IsSet(m.Name()) {
			slog.Info("initializing module", "module", m.Name())
			if err := m.Init(v); err != nil {
				slog.Error("module init failed", "module", m.Name(), "err", err)
				panic(err) // 初始化失败时，直接 panic
			}
			slog.Debug("module initialized successfully", "module", m.Name())
		} else {
			slog.Warn("module config not found, skipping init", "module", m.Name())
		}
	}
	slog.Info("all pluggable modules initialized")
}

// CloseModules 关闭所有已注册的模块
func CloseModules() {
	slog.Info("closing modules...", "count", len(modules))
	for _, m := range modules {
		slog.Debug("module close", "module", m.Name())
		if err := m.Close(); err != nil {
			slog.Error("module close failed", "module", m.Name(), "err", err)
		}
	}
}
