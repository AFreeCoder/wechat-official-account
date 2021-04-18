/*
 * @Description: Do not edit
 * @Date: 2021-04-05 22:53:40
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-18 19:50:12
 */

package logger

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var fieldKeys = []string{
	"logid",
	"status",
	"errno",
	"errmsg",
	"method",
	"path",
	"uri",
	"cost",
	"client_ip",
	"host",
	"refer",
}

// handleMap 日志字段处理方法
var handleMap = map[string]func(*gin.Context) string{
	"logid": func(c *gin.Context) string {
		return genLogid()
	},
	"status": func(c *gin.Context) string {
		return strconv.Itoa(c.Writer.Status())
	},
	"errno": func(c *gin.Context) string {
		return "0"
	},
	"errmsg": func(c *gin.Context) string {
		return c.Errors.String()
	},
	"method": func(c *gin.Context) string {
		return c.Request.Method
	},
	"path": func(c *gin.Context) string {
		return c.Request.URL.Path
	},
	"uri": func(c *gin.Context) string {
		return c.Request.URL.RequestURI()
	},
	"cost": func(c *gin.Context) string {
		return ""
	},
	"client_ip": func(c *gin.Context) string {
		return c.Request.RemoteAddr
	},
	"host": func(c *gin.Context) string {
		return c.Request.Host
	},
	"refer": func(c *gin.Context) string {
		return c.Request.Referer()
	},
}

// genLogid logid生成方式：随机数+时间戳
func genLogid() string {
	// 获取时间戳后10位
	timestamp := time.Now().Nanosecond() / 1000
	timeStr := strconv.Itoa(timestamp)

	// 生成小于100000000随机数
	rand.Seed(int64(timestamp))
	randInd := rand.Intn(1000000000)
	randStr := strconv.Itoa(randInd)

	// 拼接logid
	return randStr + timeStr
}
