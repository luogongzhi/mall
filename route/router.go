package route

import (
	"github.com/gin-gonic/gin"
	"mall/pkg"
	"mall/pkg/e"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{pkg.Code: http.StatusOK, pkg.MESSAGE: e.GetMsg(http.StatusOK)})
		})
		return r
	}
}
