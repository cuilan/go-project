//go:build windows && amd64

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

var elog debug.Log

type goService struct{}

func (s *goService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}

	// 这是服务的核心循环
	elog.Info(1, "服务启动中...")
	go runServiceLogic()

	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

loop:
	for c := range r {
		switch c.Cmd {
		case svc.Interrogate:
			changes <- c.CurrentStatus
			time.Sleep(100 * time.Millisecond)
			changes <- c.CurrentStatus
		case svc.Stop, svc.Shutdown:
			elog.Info(1, "服务停止中...")
			// 在这里执行清理工作
			break loop
		default:
			elog.Error(1, fmt.Sprintf("未知的控制请求 #%d", c))
		}
	}

	changes <- svc.Status{State: svc.StopPending}
	return
}

// 这个函数包含了应用的实际业务逻辑。
// 它在 Execute 方法内部被调用。
func runServiceLogic() {
	// 调用与非服务版本共享的应用核心逻辑。
	// 我们传入一个函数，用于在服务关闭时执行。
	err := RunApp(func() {
		// 这是优雅关闭时执行的逻辑
		elog.Info(1, "优雅关闭逻辑已执行。")
	})
	if err != nil {
		elog.Error(1, fmt.Sprintf("应用出错退出: %v", err))
	} else {
		elog.Info(1, "应用已优雅退出。")
	}
}

func main() {
	// 判断当前是否作为 Windows 服务运行
	isService, err := svc.IsWindowsService()
	if err != nil {
		fmt.Printf("无法判断是否是 Windows 服务: %v", err)
		os.Exit(1)
	}

	if !isService {
		// 在控制台或终端中运行 (用于调试)
		fmt.Println("正在以控制台模式运行。服务相关操作请使用管理脚本。")
		// 直接调用应用的核心逻辑
		err := RunApp(nil) // 控制台模式不需要特殊的关闭钩子函数
		if err != nil {
			fmt.Printf("应用运行出错: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// 作为 Windows 服务运行
	var errElog error
	elog, errElog = eventlog.Open(getExeBaseName())
	if errElog != nil {
		return // 无法记录日志，直接退出
	}
	defer elog.Close()

	elog.Info(1, "服务启动...")
	err = svc.Run(getExeBaseName(), &goService{})
	if err != nil {
		elog.Error(1, fmt.Sprintf("服务运行失败: %v", err))
		return
	}
	elog.Info(1, "服务已停止")
}

func getExeBaseName() string {
	exePath, err := os.Executable()
	if err != nil {
		return "your-go-project"
	}
	return strings.TrimSuffix(filepath.Base(exePath), ".exe")
}
