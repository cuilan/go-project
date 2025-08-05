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

// handleHealthCheck 是一个健康检查的 handler。
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
