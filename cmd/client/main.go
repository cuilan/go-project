//go:build darwin || linux

// Copyright (c) 2024 Go Project Contributors
// Licensed under the MIT License. See LICENSE file in the project root for license information.

// Package main 提供了基于Gin框架的HTTP API服务器
//
//	@title						Go Project API (Gin)
//	@version					1.0
//	@description				这是一个使用Gin框架构建的Go项目API服务器
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:8080
//	@BasePath					/
//	@schemes					http https
//	@produce					json
//	@consumes					json
package main

import "log/slog"

func main() {
	if err := RunApp(shutdownHook); err != nil {
		slog.Error("Application run error", "error", err)
	}
}

// shutdownHook 是应用关闭时的钩子函数，用于执行一些清理操作
func shutdownHook() {
	slog.Info("应用程序关闭，执行钩子...")
	slog.Info("Shutting down server...")
}
