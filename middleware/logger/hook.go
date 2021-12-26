package logger

import (
	"github.com/sirupsen/logrus"

	"main/util"
)

// Hook logrus.Hook 的实现，用于添加 request id 到当前请求中，后续所有的日志的打印都将会使用本次请求中的 request id，方便请求的全链路追踪
type Hook struct {
	RequestID string
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data[util.XRequestID] = hook.RequestID
	return nil
}
