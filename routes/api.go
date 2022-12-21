package routes

import (
	"uc/app/controller/api"
	"uc/app/middleware"

	"github.com/gin-gonic/gin"
)

func apiRoute(s *gin.Engine) {
	v := s.Group("/api")

	{
		v.Use(gin.Recovery())

		//v.GET("/user", api.Ping)
		v.POST("/login", api.Login)
		v.GET("/user", api.Ping)
		v.GET("/cache", api.Cache)
		v.GET("/config", api.Config)

		v.Use(middleware.JWT()).GET("/user/info", api.UserList)
		v.POST("/ping", api.Ping)

		v.Static("/idx", "D:\\go-programs")
	}
}
