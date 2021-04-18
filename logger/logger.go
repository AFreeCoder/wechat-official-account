/*
 * @Description: Do not edit
 * @Date: 2021-04-05 02:10:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-18 17:20:57
 */

package logger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
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

func initLogger() {
	config := defaultLogConfig

	// 指定日志路径
	logFilePath := "./logs/"
	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		panic("dir ./logs create failed!")
	}

	// 指定日志文件，暂时不做日志切分
	f, err := os.OpenFile(logFilePath+config.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic("open log file filed")
	}

	// 设置输出
	logger.SetOutput(f)

	// 设置日志级别
	if gin.Mode() == gin.ReleaseMode {
		logger.SetLevel(logger.InfoLevel)
	} else {
		logger.SetLevel(logger.DebugLevel)
	}

	// 设置日志格式
	logger.SetFormatter(&logger.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// initLogFields 初始化日志字段，这些都是必须打印的字段
func initLogFields(c *gin.Context) logger.Fields {
	var logFields = logger.Fields{}
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

// Logger 中间件
func Logger() gin.HandlerFunc {
	initLogger()
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Bind logFields to context
		logFields := initLogFields(c)
		c.Set(logFieldKey, logger.Fields{})

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()

		logFields["cost"] = end.Sub(start).Milliseconds()

		logFieldsExt := c.MustGet(logFieldKey).(logger.Fields)
		for k, v := range logFieldsExt {
			logFields[k] = v
		}
		logger.WithFields(logFields).Info("")
	}
}

// AddField 写入info日志
func AddField(c *gin.Context, key, msg string) {
	logFieldsExt := c.MustGet(logFieldKey).(logger.Fields)
	logFieldsExt[key] = msg
}

// Warn 打印warning日志
func Warn(c *gin.Context, msg string) {
	logFields := c.MustGet(logFieldKey).(logger.Fields)
	logger.WithFields(logFields).Warn("")
}
