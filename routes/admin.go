package routes

import "github.com/gin-gonic/gin"

func admin(s *gin.Engine) {
	r := s.Group("admin")

	{
		r.GET("/index", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"msg": "admin",
			})
		})
	}
}
