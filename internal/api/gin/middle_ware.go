package gin

import (
	"github.com/gin-gonic/gin"
)

// MiddleWare 全局中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "value")
		// 执行请求
		c.Next()
		c.Writer.Header().Set("content-type", "application/json")
	}
}
