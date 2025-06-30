package main

import (
	"context"
	"flag"
	"go-project/internal/conf"
	"go-project/internal/http"
	"go-project/version"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// runApp 包含应用的核心逻辑。
// 此函数被 main.go 和 main_windows.go 中的 main() 函数调用。
// shutdownHook 是一个可选的函数，用于在程序优雅退出时执行。
func runApp(shutdownHook func()) error {
	var confPath string
	// 使用 flag 来处理配置路径，这种方式更健壮。
	flag.StringVar(&confPath, "config-dir", conf.DefaultPath, "配置文件目录路径")
	flag.Parse()

	// 初始化配置
	conf.Unmarshal(confPath)

	// 打印 Banner 和版本信息
	conf.PrintBanner()
	logs.Info("================================================")
	logs.Info("|            Your Go Project Service           |")
	logs.Info("------------------------------------------------")
	logs.Info("> 操作系统: %s / 架构: %s", runtime.GOOS, runtime.GOARCH)
	logs.Info("> Go 版本: %s", version.GoVersion)
	logs.Info("> 项目版本: %s", version.Version)
	logs.Info("> 配置文件路径: %s", confPath)
	logs.Info("================================================")

	// 在一个 goroutine 中启动 HTTP 服务
	go http.Server(conf.HttpPort, conf.Mode)

	// 处理优雅退出
	quit := make(chan os.Signal, 1)
	// 监听 SIGINT 和 SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Info("服务准备关闭...")

	if shutdownHook != nil {
		shutdownHook()
	}

	// 等待5秒，让正在处理的请求完成
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logs.Info("服务已优雅退出")
	return nil
}
