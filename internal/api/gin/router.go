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

	router.GET(Root, func(c *gin.Context) {
		c.JSON(http.StatusOK, api.Success())
	})
	router.GET(Ping, func(c *gin.Context) {
		c.JSON(http.StatusOK, api.SuccessWithData("pong"))
	})
	router.GET(Health, healthCheck)

	userGroup := router.Group(UserGroup)
	{
		userGroup.GET(UserLogin, userLogin)
	}

}
