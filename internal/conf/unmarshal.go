package conf

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"

	"go-project/internal/logger"
	"go-project/internal/module"
)

const (
	ConfigType  = "yaml"      // 配置文件格式
	DefaultPath = "./configs" // 默认配置路径
	DefaultName = "app"       // 默认配置名称
)

// Unmarshal 解析配置文件
// confPath 配置文件路径
func Unmarshal(confPath string) {
	UnmarshalProfile(confPath, "")
}

// UnmarshalProfile 解析配置文件
// confPath 配置文件路径
// profile 配置文件名称
func UnmarshalProfile(confPath, profile string) {
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

	// 3. 如果 profile 存在，则加载并合并对应的环境配置文件
	if App.Profile != "" {
		v.SetConfigName(DefaultName + "-" + App.Profile)
		if err := v.MergeInConfig(); err != nil {
			// 如果 profile 配置文件不存在，可以只打印警告而不 panic，取决于您的需求
			// 在这里我们选择 panic，因为 profile 通常是必须的
			panic(fmt.Errorf("merge profile config file failed: %w", err))
		}
	}

	// 4. 将最终的配置 Unmarshal 到全局变量或结构体中
	App.Name = v.GetString("app.name")

	// 初始化日志记录器
	logger.InitLogger(
		v.GetString("log.path"),
		App.Name,
		v.GetString("log.level"),
		v.GetBool("log.enable2file"),
		v.GetInt("log.maxdays"),
	)

	// 5. 设置和初始化其他模块
	slog.Info("======== Application config ========>")
	slog.Info("app name", "name", App.Name)
	slog.Info("app profile", "profile", App.Profile)
	// slog.Info("http port", "port", App.HttpPort)
	// slog.Info("app run mode", "mode", App.Mode)
	slog.Info("-------------------------------------")
	slog.Info("log enable to file", "enabled", v.GetBool("log.enable2file"))
	slog.Info("log path", "path", v.GetString("log.path"))
	slog.Info("log level", "level", v.GetString("log.level"))
	slog.Info("log max days", "maxdays", v.GetInt("log.maxdays"))
	slog.Info("-------------------------------------")

	// 将 viper 实例传递给模块初始化函数
	initModules(v)

	slog.Info("<======= Application config =========")
}

// initModules 根据类型解析配置文件
func initModules(v *viper.Viper) {
	slog.Info("======== Initializing modules ========>")
	module.InitModules(v)
	slog.Info("<======= Modules initialized =========")
}
