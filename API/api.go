package api

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	SetupAuthen(router)
}
