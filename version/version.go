package version

import (
	"fmt"
	"runtime"
)

var (
	Major     = "1"        // 主版本号
	Minor     = "0"        // 次版本号
	Patch     = "0"        // 修订版本号
	Commit    = "8bcf8fe7" // 提交版本号
	Milestone = "Alpha"    // 里程碑版本
)

var (
	Package   = "your-go-project-name"
	Version   = fmt.Sprintf("%s.%s.%s.%s_%s", Major, Minor, Patch, Commit, Milestone)
	Revision  = ""
	GoVersion = runtime.Version()
)
