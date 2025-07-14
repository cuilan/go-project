package version

import (
	"fmt"
	"runtime"
)

var (
	major     = "0" // 主版本号
	minor     = "1" // 次版本号
	patch     = "0" // 修订版本号
	commit    = ""  // 提交版本号
	milestone = ""  // 里程碑版本
)

var (
	Package   = "your-go-project-name"
	Revision  = ""
	GoVersion = runtime.Version()
)

func Version() string {
	version := fmt.Sprintf("%s.%s.%s", major, minor, patch)
	if commit != "" {
		version += fmt.Sprintf("-%s", commit)
	}
	if milestone != "" {
		version += fmt.Sprintf("-%s", milestone)
	}
	return version
}
