package user

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/serializer"
	"mall/service"
	"net/http"
)

type IUserApi interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userApiImplementation struct{}

// NewUserApi 返回接口实现类（赋值给接口）
func NewUserApi() IUserApi {
	return &userApiImplementation{}
}

// Register 用户注册
func (*userApiImplementation) Register(c *gin.Context) {
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

// Login 用户登陆
func (*userApiImplementation) Login(c *gin.Context) {
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
