package auth

import (
	"github.com/gin-gonic/gin"

	"main/service"
	"main/util"
)

func Auth(ctx *gin.Context) {
	util.LOG.Info("controller auth: start")

	resp := service.Auth()

	util.LOG.Info("controller auth: end")

	response := util.Response{Ctx: ctx}
	response.Response(resp.Code, resp.Data, resp.Message)
}
