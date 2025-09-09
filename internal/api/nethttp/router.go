package nethttp

import (
	"encoding/json"
	"go-project/internal/api"
	"net/http"
)

// registerRoutes 注册所有路由。
func (s *Server) registerRoutes() {
	s.router.HandleFunc("/health", s.handleHealthCheck())

	s.router.HandleFunc("/user/register", s.handleUserRegister())
	s.router.HandleFunc("/user/login", s.handleUserLogin())
}

// handleHealthCheck 健康检查接口
//
//	@Summary		Health Check
//	@Description	检查服务器健康状态
//	@Tags			Common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	api.SuccessResponse{data=api.HealthResponse}	"健康检查成功"
//	@Failure		500	{object}	api.ErrorResponse									"服务器内部错误"
//	@Router			/health [get]
func (s *Server) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		result := api.SuccessWithData(map[string]string{"status": "ok"})
		// 返回 http 状态码 200
		w.WriteHeader(http.StatusOK)
		// 返回 json 数据
		json.NewEncoder(w).Encode(result)
	}
}
