/*
 * @Description: Do not edit
 * @Date: 2021-04-05 22:53:40
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-06 00:35:37
 */

package logger

import (
	"strconv"

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

var handleMap = map[string]func(*gin.Context) string{
	"logid": func(c *gin.Context) string {
		return ""
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
		return ""
	},
	"host": func(c *gin.Context) string {
		return c.Request.Host
	},
	"refer": func(c *gin.Context) string {
		return c.Request.Referer()
	},
}
