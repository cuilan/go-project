package main

import (
	"context"
	"flag"
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
	// 使用 flag 来处理配置路径，这种方式更健壮。
	flag.StringVar(&confPath, "config-dir", conf.DefaultPath, "配置文件目录路径")
	flag.Parse()

	// 初始化配置
	conf.Unmarshal(confPath)

	// 打印 Banner 和版本信息
	if conf.App.Profile == "dev" {
		conf.PrintGoBanner()
	} else {
		conf.PrintBanner()
	}
	slog.Info("================================================")
	slog.Info("|            Your Go Project Service           |")
	slog.Info("------------------------------------------------")
	slog.Info("> OS / Arch", "os", runtime.GOOS, "arch", runtime.GOARCH)
	slog.Info("> Go Version", "version", version.GoVersion)
	slog.Info("> Project Version", "version", version.Version())
	slog.Info("> Config Path", "path", confPath)
	slog.Info("================================================")

	// 在一个 goroutine 中启动 HTTP 服务
	// go http.Server(conf.HttpPort, conf.Mode)
	rdb.RedisClient.Set(context.Background(), "name", "zhangyan", 0*time.Second)
	value := rdb.RedisClient.Get(context.Background(), "name")
	slog.Info("redis value", "value", value)

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

	// close http server
	// http.Stop()

	// close all modules
	module.CloseModules()

	slog.Info("server exited gracefully")
	return nil
}
