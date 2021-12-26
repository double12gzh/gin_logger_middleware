package visit

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"net/http"

	"main/service"
	"main/util"
)

func Visit(ctx *gin.Context) {
	// 打印日志时默认添加当前请求中使用的 request id。如果请求中没有传，则使用自生成的
	util.LOG.Info("controller visit: start")

	resp := service.Visit()

	util.LOG.Info("controller visit: end")

	response := util.Response{Ctx: ctx}
	response.Response(resp.Code, resp.Data, resp.Message)
}

func DO(ctx *gin.Context) {
	var u model.Visitor

	response := util.Response{Ctx: ctx}
	if err := ctx.ShouldBindJSON(&u); err == nil {
		resp := service.Do()
		response.Response(resp.Code, resp.Data, resp.Message)
	} else {
		response.Response(http.StatusBadRequest, nil, err.Error())
	}
}
