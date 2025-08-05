package gin

import (
	"go-project/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, api.SuccessWithMsg("health check success"))
}
