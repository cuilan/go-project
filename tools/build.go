//go:build tools

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// main 是构建脚本的入口点。
func main() {
	log.Println("开始执行跨平台构建...")

	// 从环境变量或默认值获取平台列表
	platforms := getPlatforms()
	appNames := getAppNames()
	version := getVersion()
	commit := getCommit()

	// 创建 bin 目录
	binDir, err := createBuildDir("bin")
	if err != nil {
		log.Fatalf("创建 bin 目录失败: %v", err)
	}

	var wg sync.WaitGroup
	for _, appName := range appNames {
		for _, p := range platforms {
			wg.Add(1)
			go buildForPlatform(p, appName, version, commit, binDir, &wg)
		}
	}
	wg.Wait()

	log.Println("所有构建任务已完成！")
}

// buildForPlatform 为指定的平台执行构建。
func buildForPlatform(platform, appName, version, commit, buildDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	parts := strings.Split(platform, "/")
	if len(parts) != 2 {
		log.Printf("警告: 无效的平台格式 '%s'，已跳过。", platform)
		return
	}
	goos, goarch := parts[0], parts[1]

	// 构造输出文件名
	var outputName string
	if commit != "" {
		outputName = fmt.Sprintf("%s_%s_%s_%s_%s", appName, version, commit, goos, goarch)
	} else {
		outputName = fmt.Sprintf("%s_%s_%s_%s", appName, version, goos, goarch)
	}
	if goos == "windows" {
		outputName += ".exe"
	}
	outputPath := filepath.Join(buildDir, outputName)

	log.Printf("正在为 %s/%s 构建: %s", goos, goarch, outputName)

	// 设置 go build 命令
	cmd := exec.Command("go", "build", "-o", outputPath, "./cmd/"+appName)
	cmd.Env = append(os.Environ(),
		"GOOS="+goos,
		"GOARCH="+goarch,
		"CGO_ENABLED=0", // 默认禁用 CGO
	)

	// Windows 平台通常需要 CGO 来支持某些特性
	if goos == "windows" {
		cmd.Env = append(cmd.Env, "CGO_ENABLED=1")
	}

	// 执行命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("错误: 为 %s/%s 构建失败。\n输出: %s\n错误: %v", goos, goarch, string(output), err)
	} else {
		log.Printf("成功: 已为 %s/%s 完成构建。", goos, goarch)
	}
}

// getPlatforms 从环境变量 "PLATFORMS" 中获取平台列表，如果环境变量未设置，则返回默认列表。
func getPlatforms() []string {
	if p := os.Getenv("PLATFORMS"); p != "" {
		return strings.Split(p, " ")
	}
	// 默认当前平台
	return []string{runtime.GOOS + "/" + runtime.GOARCH}
}

// getAppNames 从环境变量 "COMMANDS" 中获取应用名称列表，如果未设置，则返回默认值。
func getAppNames() []string {
	if names := os.Getenv("COMMANDS"); names != "" {
		return strings.Split(names, " ")
	}
	return []string{"your-app"}
}

// getVersion 从环境变量 "VERSION" 中获取版本号，如果未设置，则尝试从 git 获取。
func getVersion() string {
	if v := os.Getenv("VERSION"); v != "" {
		return v
	}
	// 尝试从 git describe 获取版本
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	out, err := cmd.Output()
	if err == nil {
		return strings.TrimSpace(strings.TrimPrefix(string(out), "v"))
	}
	return "0.1.0"
}

// getCommit 从环境变量 "COMMIT" 中获取 git commit hash，如果未设置，则尝试从 git 获取。
func getCommit() string {
	if c := os.Getenv("COMMIT"); c != "" {
		return c
	}
	return ""
}

// createBuildDir 创建用于存放构建产物的目录。
func createBuildDir(dirName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	buildDir := filepath.Join(wd, dirName)
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return "", err
	}
	return buildDir, nil
}
