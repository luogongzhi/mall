package user

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"mall/service"
	"net/http"
)

type IUserAddressApi interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type userAddressApiImplementation struct{}

func NewUserAddressApi() IUserAddressApi {
	return &userAddressApiImplementation{}
}

// List 根据Id查询用户地址信息
func (*userAddressApiImplementation) List(c *gin.Context) {
	var userAddressService service.UserAddressService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := userAddressService.List(c.Request.Context(), claims.Id)
	c.JSON(http.StatusOK, res)
}

// Create 用户地址添加
func (*userAddressApiImplementation) Create(c *gin.Context) {
	var userAddressService service.UserAddressService
	var dto serializer.UserAddressCreateDTO
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&dto); err == nil {
		res := userAddressService.Create(c.Request.Context(), dto, claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

// Delete 用户地址删除
func (*userAddressApiImplementation) Delete(c *gin.Context) {
	var userAddressService service.UserAddressService
	var dto serializer.UserAddressDeleteDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := userAddressService.Delete(c.Request.Context(), dto.Id, claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

// Update 用户基本信息修改
func (*userAddressApiImplementation) Update(c *gin.Context) {
	var userAddressService service.UserAddressService
	var dto serializer.UserAddressUpdateDTO
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBindJSON(&dto); err == nil {
		res := userAddressService.Update(c.Request.Context(), dto, claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}
