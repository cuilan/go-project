package gin

import (
	"go-project/internal/api"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userLogin(c *gin.Context) {
	var userLoginModel UserRegisterModel
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
