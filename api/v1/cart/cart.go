package cart

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/utils"
	"mall/service"
	"net/http"
)

type ICartApi interface {
	Detail(c *gin.Context)
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
