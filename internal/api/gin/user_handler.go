package gin

import (
	"go-project/internal/api"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// userLogin 用户登录接口
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
func userLogin(c *gin.Context) {
	var userLoginModel api.UserLoginRequest
	if err := c.ShouldBind(&userLoginModel); err != nil {
		c.JSON(http.StatusBadRequest, api.FailWithMsg(err.Error()))
		return
	}

	userService := service.GetUserService()
	user, err := userService.UserLogin(c.Request.Context(), userLoginModel.Username, userLoginModel.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.FailWithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, api.SuccessWithData(user))
}
