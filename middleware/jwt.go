package middleware

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/pkg/utils"
	"mall/serializer"
	"net/http"
	"time"
)

// JWT jwt中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuthInsufficientAuthority
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt.Time.Unix() {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != http.StatusOK {
			c.Abort()
			c.JSON(http.StatusOK, serializer.ResponseResult{
				Code: code,
				Msg:  e.GetMsg(code),
				Data: nil,
			})
			return
		}
		c.Next()
	}
}
