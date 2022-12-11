package v1

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/serializer"
	"mall/service"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBindJSON(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code: e.InvalidParams,
			Msg:  e.GetMsg(e.InvalidParams),
		})
	}
}
