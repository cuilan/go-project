//go:build (darwin && amd64) || (linux && amd64)

package main

import (
	"github.com/beego/beego/v2/core/logs"
	"go-project/internal/conf"
	"go-project/internal/http"
	"go-project/version"
	"os"
	"runtime"
)

func init() {
	var confPath string
	if len(os.Args) == 1 {
		confPath = conf.DefaultPath
	} else {
		confPath = os.Args[1]
	}

	conf.Unmarshal(confPath)
}

func main() {
	conf.PrintBanner()
	logs.Info("================================================")
	logs.Info("|           大地新亚卫星云图监控程序            |")
	logs.Info("|               Linux systemd                 |")
	logs.Info("------------------------------------------------")
	logs.Info("> OS: %s / Arch: %s", runtime.GOOS, runtime.GOARCH)
	logs.Info("> Package: %s", version.Package)
	logs.Info("> GoVersion: %s", version.GoVersion)
	logs.Info("> Version: %s", version.Version)
	logs.Info("> os.Args: %s", os.Args)
	logs.Info("================================================")

	go http.Server(conf.HttpPort, conf.Mode)

	select {}
}
