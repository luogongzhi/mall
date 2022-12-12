package v1

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/serializer"
	"mall/service"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBindJSON(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBindJSON(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}
