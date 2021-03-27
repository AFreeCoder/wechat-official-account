/*
 * @Description: Do not edit
 * @Date: 2021-03-20 12:43:38
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-03-20 21:24:00
 */

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BaseController controller类型
type BaseController struct{}

// DefaultController 单例
var DefaultController = &BaseController{}

// ResponseString 返回string
func (b *BaseController) ResponseString(c *gin.Context, data string) {
	c.String(http.StatusOK, data)
	return
}

// ResponseXML 返回XML格式
func (b *BaseController) ResponseXML(c *gin.Context, m interface{}) {
	c.XML(200, m)
	return
}
