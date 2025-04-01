package http

import (
	"github.com/gin-gonic/gin"
)

func testController(c *gin.Context) {
	Success(c, "hello")
}
