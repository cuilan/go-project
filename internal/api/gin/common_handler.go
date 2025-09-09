package gin

import (
	"go-project/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthCheck 健康检查接口
//
//	@Summary		Health Check
//	@Description	检查服务器健康状态
//	@Tags			Common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	api.SuccessResponse{data=api.HealthResponse}	"健康检查成功"
//	@Failure		500	{object}	api.ErrorResponse					"服务器内部错误"
//	@Router			/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, api.SuccessWithMsg("health check success"))
}
