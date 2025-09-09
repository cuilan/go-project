package gin

import (
	"go-project/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Root   = "/"
	Ping   = "/ping"
	Health = "/health"

	UserGroup = "/user" // 配置用户组
	UserLogin = "/login"
)

func uri(router *gin.Engine) {

	// rootHandler 根路径处理器
	//
	//	@Summary		Root Endpoint
	//	@Description	API根路径，返回基本成功响应
	//	@Tags			Common
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	api.SuccessResponse	"请求成功"
	//	@Router			/ [get]
	router.GET(Root, func(c *gin.Context) {
		c.JSON(http.StatusOK, api.Success())
	})

	// pingHandler ping接口
	//
	//	@Summary		Ping
	//	@Description	Ping接口，用于测试服务器连通性
	//	@Tags			Common
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{object}	api.SuccessResponse{data=string}	"返回pong"
	//	@Router			/ping [get]
	router.GET(Ping, func(c *gin.Context) {
		c.JSON(http.StatusOK, api.SuccessWithData("pong"))
	})

	router.GET(Health, healthCheck)

	userGroup := router.Group(UserGroup)
	{
		userGroup.POST(UserLogin, userLogin)
	}

}
