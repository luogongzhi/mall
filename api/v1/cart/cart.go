package cart

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"mall/service"
	"net/http"
)

type ICartApi interface {
	Detail(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type cartApiImplementation struct{}

func NewCartApi() ICartApi {
	return &cartApiImplementation{}
}

func (*cartApiImplementation) Detail(c *gin.Context) {
	var cartService service.CartService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := cartService.Detail(c.Request.Context(), claims.Id)
	c.JSON(http.StatusOK, res)
}

func (i *cartApiImplementation) Create(c *gin.Context) {
	var cartService service.CartService
	var dto serializer.CartCreateDeleteDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := cartService.Create(c.Request.Context(), dto, claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

func (i *cartApiImplementation) Delete(c *gin.Context) {
	var cartService service.CartService
	var dto serializer.CartCreateDeleteDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := cartService.Delete(c.Request.Context(), dto, claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}
