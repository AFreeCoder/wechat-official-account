/*
 * @Description: Do not edit
 * @Date: 2021-04-05 02:10:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-07 01:30:10
 */

package logger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	logFieldKey = "logrus_log_fields_key"
)

type logConfig struct {
	FileName string
}

var defaultLogConfig = logConfig{
	FileName: "service.log",
}

func initLogger() *logrus.Logger {
	config := defaultLogConfig

	// 指定日志路径
	logFilePath := "./logs/"
	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		panic("dir ./logs create failed!")
	}

	// 指定日志文件
	f, err := os.OpenFile(logFilePath+config.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic("open log file filed")
	}
	logger := logrus.New()

	// 设置输出
	logger.SetOutput(f)

	// 设置日志级别
	if gin.Mode() == gin.ReleaseMode {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(logrus.DebugLevel)
	}

	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func initLogFields(c *gin.Context) logrus.Fields {
	var logFields = logrus.Fields{}
	for _, k := range fieldKeys {
		f, ok := handleMap[k]
		if ok {
			logFields[k] = f(c)
		} else {
			logFields[k] = ""
		}
	}
	return logFields
}

func Logger() gin.HandlerFunc {
	logger := initLogger()
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Bind logFields to context
		logFields := initLogFields(c)
		c.Set(logFieldKey, logrus.Fields{})

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()

		logFields["cost"] = end.Sub(start).Milliseconds()

		logFieldsExt := c.MustGet(logFieldKey).(logrus.Fields)
		for k, v := range logFieldsExt {
			logFields[k] = v
		}

		logger.WithFields(logFields).Info("")
	}
}

func AddLogField(c *gin.Context, key, msg string) {
	logFieldsExt := c.MustGet(logFieldKey).(logrus.Fields)
	logFieldsExt[key] = msg
}
