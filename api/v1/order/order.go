package order

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"mall/service"
	"net/http"
)

type IOrderApi interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
}

type orderApiImplementation struct{}

func NewProductApi() IOrderApi {
	return &orderApiImplementation{}
}

// Create 创建订单
func (*orderApiImplementation) Create(c *gin.Context) {
	var OrderService service.OrderService
	var dto serializer.OrderCreateDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := OrderService.Create(c.Request.Context(), claims.Id, dto)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

// Update 修改订单
func (*orderApiImplementation) Update(c *gin.Context) {
	var OrderService service.OrderService
	var dto serializer.OrderUpdateDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := OrderService.Update(c.Request.Context(), claims.Id, dto)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

// List 用户订单列
func (*orderApiImplementation) List(c *gin.Context) {
	var OrderService service.OrderService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := OrderService.List(c.Request.Context(), claims.Id)
	c.JSON(http.StatusOK, res)
}

// Delete 取消订单
func (*orderApiImplementation) Delete(c *gin.Context) {
	var OrderService service.OrderService
	var dto serializer.OrderDeleteDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := OrderService.Delete(c.Request.Context(), claims.Id, dto.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}
