package main

import (
	"go-project/internal/api/gin"
	"go-project/internal/module"
	"go-project/internal/orm/gorm"
	"go-project/internal/rdb"
)

// InjectModules 是应用中所有可插拔模块的列表。
// 当需要添加新模块时，请在此处调用其 NewModule() 函数。
// 务必确保模块的 NewModule() 函数无参数，
// 且返回的 module.Module 实现了 module.Module 接口。
var InjectModules = []module.Module{
	// example: new_module.NewModule(),
	rdb.NewModule(),

	// 数据库模块 - 根据配置文件自动选择
	// 如果配置文件中同时存在 gorm 和 gosql 配置，优先使用 gorm
	// 如果只存在其中一个配置，则使用对应的模块
	// gosql.NewModule(), // database/sql 实现
	gorm.NewModule(), // GORM 实现

	// nethttp.NewModule(), // net/http 实现
	gin.NewModule(), // gin 实现
}
