package util

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LOG = logrus.New()

func init() {
	Logger()
}

func Logger() {
	var logFilePath string
	startTime := time.Now()

	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}

	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		fmt.Println(err.Error())
	}

	logFileName := startTime.Format("2006-01-02 15:04:05") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	//打开文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("write file log error", err)
	}

	//实例化
	LOG.Out = src
	LOG.SetLevel(logrus.DebugLevel)
	LOG.SetReportCaller(true)
	LOG.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:         "[MSG]",
			logrus.FieldKeyLevel:       "[LEVEL]",
			logrus.FieldKeyFunc:        "[FUNC]",
			logrus.FieldKeyTime:        "[TIME]",
			logrus.FieldKeyFile:        "[FILE]",
			logrus.FieldKeyLogrusError: "[ERROR]",
		},
	})
}
