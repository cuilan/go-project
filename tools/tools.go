//go:build tools

// tools 包用于追踪 go.mod 文件中的构建时依赖项。
// 更多信息请参考: https://go.dev/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "honnef.co/go/tools/cmd/staticcheck"
)
