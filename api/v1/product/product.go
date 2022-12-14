package product

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/serializer"
	"mall/service"
	"net/http"
	"strconv"
)

type IProductApi interface {
	Detail(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
}

type productApiImplementation struct{}

func NewProductApi() IProductApi {
	return &productApiImplementation{}
}

func (*productApiImplementation) Detail(c *gin.Context) {
	var productService service.ProductService
	if value, ok := c.GetQuery("id"); ok == true {
		id, _ := strconv.ParseInt(value, 10, 64)
		res := productService.Detail(c.Request.Context(), uint64(id))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

func (*productApiImplementation) Create(c *gin.Context) {
	var productService service.ProductService
	var dto serializer.ProductCreateDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		res := productService.Create(c.Request.Context(), dto)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

func (*productApiImplementation) Delete(c *gin.Context) {
	var productService service.ProductService
	var dto serializer.ProductDeleteDTO
	if err := c.ShouldBindJSON(&dto); err == nil {
		res := productService.Delete(c.Request.Context(), dto.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}

func (*productApiImplementation) Update(c *gin.Context) {
	var productService service.ProductService
	var dto serializer.ProductUpdateDTO
	_ = c.ShouldBindJSON(&dto)
	res := productService.Update(c.Request.Context(), dto)
	c.JSON(http.StatusOK, res)
}

func (*productApiImplementation) List(c *gin.Context) {
	var productService service.ProductService
	var dto serializer.PaginateDTO
	if err := c.ShouldBindQuery(&dto); err == nil {
		res := productService.List(c.Request.Context(), dto)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.ResponseResult{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}
