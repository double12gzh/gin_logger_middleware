package v1

import (
	"github.com/gin-gonic/gin"

	ctrlAuth "main/controller/auth"
)

func SetupAuthorized(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	auth.GET("", ctrlAuth.Auth)
}
