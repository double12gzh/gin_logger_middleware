package util

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type InnerResponse struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	StartTime string      `json:"start_time"`
	EndTime   string      `json:"end_time"`
}

func (r *Response) Response(code int, data interface{}, message string) {
	reqID := r.Ctx.GetString(XRequestID)

	r.Ctx.JSON(code, InnerResponse{
		Code:      code,
		Data:      data,
		Message:   fmt.Sprintf("%s. request_id: %s.", message, reqID),
		RequestID: reqID,
		StartTime: fmt.Sprintf("%v", r.Ctx.GetString(XStartTime)),
		EndTime:   fmt.Sprintf("%v", time.Now().Unix()),
	})

	return
}
