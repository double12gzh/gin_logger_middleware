package apis

import (
	"github.com/gin-gonic/gin"

	v1 "main/apis/v1"
	"main/middleware/logger"
)

func SetupRouters(r *gin.Engine) {
	{
		r.Use(logger.Tracing())
	}

	v1Router := r.Group("v1")

	v1.SetupAuthorized(v1Router)
	v1.SetupVisit(v1Router)
}
