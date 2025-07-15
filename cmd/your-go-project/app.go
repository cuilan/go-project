package main

import (
	"context"
	"flag"
	"fmt"
	"go-project/internal/conf"
	"go-project/internal/module"
	"go-project/internal/rdb"
	"go-project/version"
	"log/slog"
	"runtime"

	"os"
	"os/signal"
	"syscall"
	"time"
)

// RunApp 包含应用的核心逻辑。
// 此函数被 main.go 和 main_windows.go 中的 main() 函数调用。
// shutdownHook 是一个可选的函数，用于在程序优雅退出时执行。
func RunApp(shutdownHook func()) error {
	var confPath string
	var showVersion bool
	flag.StringVar(&confPath, "config-dir", conf.DefaultPath, "配置文件目录路径")
	flag.BoolVar(&showVersion, "version", false, "显示版本信息")
	flag.Parse()

	if showVersion {
		fmt.Println(version.Version())
		return nil
	}

	// --- 配置加载 ---
	conf.UnmarshalProfile(confPath, conf.App.Profile, InjectModules)

	// --- 打印 Banner ---
	if conf.App.Profile == "dev" {
		conf.PrintGoBanner()
	} else {
		conf.PrintBanner()
	}
	slog.Info("================================================")
	slog.Info("|            Your Go Project Service           |")
	slog.Info("------------------------------------------------")
	slog.Info(">", "OS", runtime.GOOS, "Arch", runtime.GOARCH)
	slog.Info("> Go", "Version", version.GoVersion)
	slog.Info("> Project", "Version", version.Version())
	slog.Info("> Config", "Path", confPath)
	slog.Info("================================================")

	// ===== 业务逻辑写在下面 =====

	rdb.GetRedis().Set(context.Background(), "name", "zhangyan", 0*time.Second)
	value := rdb.GetRedis().Get(context.Background(), "name")
	slog.Debug("debug value", "value", value)
	slog.Info("info value", "value", value)
	slog.Warn("warn value", "value", value)
	slog.Error("error value", "value", value)

	// ===== 业务逻辑写在上面 =====

	// 处理优雅退出
	quit := make(chan os.Signal, 1)
	// 监听 SIGINT 和 SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("shutting down server...")

	if shutdownHook != nil {
		shutdownHook()
	}

	// 等待5秒，让正在处理的请求完成
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭所有模块
	module.CloseModules()

	slog.Info("server exited gracefully")
	slog.Info("Goodbye!")
	return nil
}
