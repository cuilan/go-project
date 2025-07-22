package main

import (
	"go-project/internal/module"
	"go-project/internal/rdb"
)

// InjectModules 是应用中所有可插拔模块的列表。
// 当需要添加新模块时，请在此处调用其 NewModule() 函数。
// 务必确保模块的 NewModule() 函数无参数，
// 且返回的 module.Module 实现了 module.Module 接口。
var InjectModules = []module.Module{
	rdb.NewModule(),
	// example: new_module.NewModule(),
}
