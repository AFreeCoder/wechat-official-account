/*
 * @Description: Do not edit
 * @Date: 2021-03-20 01:40:57
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-03-21 00:49:40
 */

package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/wechat-offical-account/httpserver/controller"
)

// RouterSetup 路由设置
func RouterSetup(r *gin.Engine) {
	r.GET("/wechat", controller.DefaultController.Get)
	r.POST("/wechat", controller.DefaultController.Post)
}
