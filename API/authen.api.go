package api

import (
	"github.com/gin-gonic/gin"
)

func SetupAuthen(router *gin.Engine) {
	authenAPI := router.Group("/api")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{"result": "login"})
}

func register(c *gin.Context) {
	c.JSON(200, gin.H{"result": "register"})
}
