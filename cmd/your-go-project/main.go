//go:build (darwin && amd64) || (linux && amd64) || (darwin && arm64)

package main

import "log/slog"

func main() {
	if err := RunApp(shutdownHook); err != nil {
		slog.Error("Application run error", "error", err)
	}
}

// shutdownHook 是应用关闭时的钩子函数，用于执行一些清理操作
func shutdownHook() {
	slog.Info("应用程序关闭，发送消息...")
	slog.Info("Shutting down server...")
}
