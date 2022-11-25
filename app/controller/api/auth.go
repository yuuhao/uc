package api

import (
	"fmt"
	"net/http"
	"uc/utils"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var demoUser = Auth{
	Username: "demo",
	Password: "123123",
}

func Login(c *gin.Context) {
	auth := Auth{}
	data := make(map[string]interface{})
	if err := c.ShouldBind(&auth); err != nil {
		c.String(http.StatusOK, `password is fail 1`)
		return
	}
	if !valid(auth) {
		c.String(http.StatusOK, `password is fail 2`)
		return
	}
	token, err := utils.GenerateToken(auth.Username, auth.Password)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("%+v", err))
		return
	}
	data["token"] = token
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": data,
	})
}

// struct 可比较吗， 答：
func valid(auth Auth) bool {
	if auth.Username != demoUser.Username || auth.Password != demoUser.Password {
		return false
	}
	return true
}
