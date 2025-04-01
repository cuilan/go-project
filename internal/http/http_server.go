package http

import (
	"github.com/gin-gonic/gin"
	"go-project/internal/conf"
)

const (
	Root      = "/"       // App名称
	PushGroup = "/push"   // 配置组
	Health    = "/health" // 健康检查

	TestController = "/getName"
)

func Server(httpPort, mode string) {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if mode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 创建路由
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())

	// 绑定路由规则，执行的函数， gin.Context，封装了request和response
	r.GET(Root, func(c *gin.Context) {
		Success(c, conf.AppName)
	})

	// push config group
	pushGroup := r.Group(PushGroup)
	{
		pushGroup.GET(Health, healthCheck)
		pushGroup.GET(TestController, testController)
	}

	err := r.Run(":" + httpPort)
	if err != nil {
		return
	}
}

func healthCheck(c *gin.Context) {
	resp := HealthResp{
		Satellite: conf.Type,
	}
	Success(c, &resp)
}
