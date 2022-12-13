package route

import (
	"github.com/gin-gonic/gin"
	api "mall/api/v1"
	"mall/middleware"
	_ "mall/middleware"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	registry := &api.Registry{}
	registry.NewRegister()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, serializer.ResponseResult{
				Code: http.StatusOK,
				Msg:  e.GetMsg(http.StatusOK),
			})
		})

		v1.POST("/user.register", registry.UserApi.Register)
		v1.POST("/user.login", registry.UserApi.Login)

		// 需要登录
		authed := v1.Group("")
		authed.Use(middleware.JWT())
		{
			// user模块
			authed.GET("/user.list", registry.UserApi.List)
			authed.POST("/user.update", registry.UserApi.Update)
			authed.GET("/user_address.list", registry.UserAddressApi.List)
			authed.POST("/user_address.update", registry.UserAddressApi.Update)
			authed.POST("/user_address.create", registry.UserAddressApi.Create)
			authed.POST("/user_address.delete", registry.UserAddressApi.Delete)
		}
	}
	return r
}
