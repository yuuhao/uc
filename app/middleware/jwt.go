package middleware

import (
	"net/http"
	"strconv"
	"time"
	"uc/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 1
		token := c.GetHeader("token")
		if token == "" {
			code = 0
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 2
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 3
			}
		}
		if code != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "认证失败" + strconv.Itoa(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
