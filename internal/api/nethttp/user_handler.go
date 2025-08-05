package nethttp

import (
	"encoding/json"
	"go-project/internal/api"
	"go-project/internal/service"
	"net/http"
)

func (s *Server) handleUserRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(api.FailWithMsg("invalid request body: " + err.Error()))
			return
		}
		defer r.Body.Close()

		if req.Username == "" || req.Password == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(api.FailWithMsg("username or password is empty"))
			return
		}

		userService := service.GetUserService()
		err := userService.UserRegister(r.Context(), req.Username, req.Password)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(api.FailWithMsg(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		result := api.SuccessWithData(map[string]string{"status": "ok"})
		json.NewEncoder(w).Encode(result)
	}
}

func (s *Server) handleUserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(api.FailWithMsg("invalid request body: " + err.Error()))
			return
		}
		defer r.Body.Close()

		if req.Username == "" || req.Password == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(api.FailWithMsg("username or password is empty"))
			return
		}

		userService := service.GetUserService()
		user, err := userService.UserLogin(r.Context(), req.Username, req.Password)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(api.FailWithMsg(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// 登录成功后返回用户信息
		result := api.SuccessWithData(user)
		json.NewEncoder(w).Encode(result)
	}
}
