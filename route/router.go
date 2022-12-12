package route

import (
	"github.com/gin-gonic/gin"
	api "mall/api/v1"
	"mall/middleware"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, serializer.ResponseResult{
				Code: http.StatusOK,
				Msg:  e.GetMsg(http.StatusOK),
			})
		})

		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 需要登录
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// user模块
			_ = authed.Group("/user")
			{
			}
		}
	}
	return r
}
