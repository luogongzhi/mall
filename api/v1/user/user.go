package user

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"mall/service"
	"net/http"
)

type IUserApi interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Detail(c *gin.Context)
	Update(c *gin.Context)
}

type userApiImplementation struct{}

// NewUserApi 返回接口实现类（赋值给接口）
func NewUserApi() IUserApi {
	return &userApiImplementation{}
}

// Register 用户注册
func (*userApiImplementation) Register(c *gin.Context) {
	var userService service.UserService
	var dto serializer.UserLoginRegisterDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		res := userService.Register(c.Request.Context(), dto)
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
	var userService service.UserService
	var dto serializer.UserLoginRegisterDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		res := userService.Login(c.Request.Context(), dto)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

// Detail 根据Id查询用户基本信息
func (*userApiImplementation) Detail(c *gin.Context) {
	var userService service.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := userService.Detail(c.Request.Context(), claims.Id)
	c.JSON(http.StatusOK, res)
}

// Update 用户基本信息修改
func (*userApiImplementation) Update(c *gin.Context) {
	var userService service.UserService
	var dto serializer.UserUpdateDTO
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&dto); err == nil {
		res := userService.Update(c.Request.Context(), dto, claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}
