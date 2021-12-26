package util

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	XRequestID = "X-Request-ID"
	XStartTime = "X-Start-Time"
)

var (
	random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

func GetIDFromRequest(c *gin.Context) string {
	if v, ok := c.Get(XRequestID); ok {
		if requestID, ok := v.(string); ok {
			return requestID
		}
	}

	if reqID := c.Request.Header.Get(XRequestID); reqID != "" {
		return reqID
	}

	return GenRequestID()
}

func GenRequestID() string {
	return genUUID(16)
}

func genUUID(length int) string {
	str := func(len int) string {
		value := make([]byte, len)
		random.Read(value)

		return base64.StdEncoding.EncodeToString(value)[:len]
	}(length)

	return str
}
