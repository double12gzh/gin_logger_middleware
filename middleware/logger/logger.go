package logger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"main/util"
	"time"

	"github.com/gin-gonic/gin"
)

// Tracing 中间件。记录请求的开始及结束时间，并打印每次请求的 body 及 query
func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()
		c.Set(util.XStartTime, fmt.Sprintf("%v", startTime.Unix()))

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		requestID := util.GetIDFromRequest(c)
		util.LOG.Hooks.Add(&Hook{requestID})

		reqURI := c.Request.URL.RequestURI()
		util.LOG.Infof("start to handle request: %s %s, body: %s", c.Request.Method, reqURI, string(body))

		c.Request.Header.Set(util.XRequestID, requestID)
		c.Set(util.XRequestID, requestID)

		c.Next()

		util.LOG.Infof("finish handling request: %s %s, latency: %s", c.Request.Method, reqURI, time.Since(startTime))
	}
}
