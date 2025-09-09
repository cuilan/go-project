package nethttp

import (
	"encoding/json"
	"go-project/internal/api"
	"go-project/internal/service"
	"net/http"
)

// handleUserRegister 用户注册接口
//
//	@Summary		User Register
//	@Description	用户注册，创建新用户账号
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		api.UserRegisterRequest	true	"用户注册信息"
//	@Success		200		{object}	api.SuccessResponse{data=object}			"注册成功"
//	@Failure		400		{object}	api.ErrorResponse							"请求参数错误"
//	@Failure		409		{object}	api.ErrorResponse							"用户已存在"
//	@Failure		500		{object}	api.ErrorResponse							"服务器内部错误"
//	@Router			/user/register [post]
func (s *Server) handleUserRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.UserRegisterRequest

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

// handleUserLogin 用户登录接口
//
//	@Summary		User Login
//	@Description	用户登录，验证用户名和密码
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		api.UserLoginRequest	true	"用户登录信息"
//	@Success		200		{object}	api.SuccessResponse{data=api.User}				"登录成功，返回用户信息"
//	@Failure		400		{object}	api.ErrorResponse							"请求参数错误"
//	@Failure		401		{object}	api.ErrorResponse							"用户名或密码错误"
//	@Failure		500		{object}	api.ErrorResponse							"服务器内部错误"
//	@Router			/user/login [post]
func (s *Server) handleUserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req api.UserLoginRequest

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
