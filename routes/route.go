package routes

import "github.com/gin-gonic/gin"

func Route(e *gin.Engine) {
	apiRoute(e)
	admin(e)
}
