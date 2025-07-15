package conf

import (
	"fmt"
	"go-project/internal/logger"
	"go-project/internal/module"
	"log/slog"

	"github.com/spf13/viper"
)

const (
	ConfigType  = "yaml"      // 配置文件格式
	DefaultPath = "./configs" // 默认配置路径
	DefaultName = "app"       // 默认配置名称
)

// AppConfig 应用配置
type AppConfig struct {
	Name    string `mapstructure:"name"`
	Profile string `mapstructure:"profile"`
}

// GlobalVar 全局变量
var App = new(AppConfig)

// Unmarshal 解析配置文件
// confPath 配置文件路径
func Unmarshal(confPath string, modules []module.Module) {
	UnmarshalProfile(confPath, "", modules)
}

// UnmarshalProfile 解析配置文件
// confPath 配置文件路径
// profile 配置文件名称
func UnmarshalProfile(confPath, profile string, modules []module.Module) {
	// 1. 设置并读取基础配置文件 (app.yaml)
	v := viper.New()
	v.SetConfigName(DefaultName)
	v.SetConfigType(ConfigType)
	if confPath != "" {
		v.AddConfigPath(confPath)
	} else {
		v.AddConfigPath(DefaultPath)
	}

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read base config file failed: %w", err))
	}

	// 2. 获取 profile 设置
	// 优先使用函数传入的 profile 参数，如果为空，则尝试从配置文件中获取
	activeProfile := profile
	if activeProfile == "" {
		activeProfile = v.GetString("app.profile")
	}

	// 将最终确定的 profile 保存到全局变量
	App.Profile = activeProfile
	App.Name = v.GetString("app.name")

	// 3. 如果 profile 存在，则加载并合并对应的环境配置文件
	if App.Profile != "" {
		v.SetConfigName(DefaultName + "-" + App.Profile)
		if err := v.MergeInConfig(); err != nil {
			panic(fmt.Errorf("merge profile config file failed: %w", err))
		}
	}

	// 4. 日志模块必须提前初始化，确保日志格式统一，如果没有日志配置，按照默认配置初始化
	if v.IsSet("log") {
		var logCfg logger.LoggerConfig
		if err := v.Sub("log").Unmarshal(&logCfg); err != nil {
			panic(fmt.Errorf("unmarshal log config failed: %w", err))
		}
		// 手动初始化日志模块
		logger.InitLogger(&logCfg)
		slog.Info("logger module initialized as a core service")
	}

	// 5. 初始化其他模块
	initModules(v, modules)

	slog.Info("application config completed", "name", App.Name, "profile", App.Profile)
}

// initModules 根据类型解析配置文件
func initModules(v *viper.Viper, modules []module.Module) {
	// 显式注册传入的模块
	for _, m := range modules {
		module.Register(m)
	}
	module.InitModules(v)
}
