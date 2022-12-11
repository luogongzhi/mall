package route

import (
	"github.com/gin-gonic/gin"
	api "mall/api/v1"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, serializer.Response{
				Code: http.StatusOK,
				Msg:  e.GetMsg(http.StatusOK),
			})
		})

		// user
		user := v1.Group("/user")
		{
			user.POST("/register", api.UserRegister)
		}
	}
	return r
}
