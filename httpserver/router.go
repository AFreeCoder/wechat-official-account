/*
 * @Description: Do not edit
 * @Date: 2021-03-20 01:40:57
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-08-12 19:44:40
 */

package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wechat-official-account/httpserver/controller"
	"github.com/wechat-official-account/httpserver/controller/proxy"
)

// RouterSetup 路由设置
func RouterSetup(r *gin.Engine) {
	r.GET("/wechat", controller.DefaultController.Get)
	r.GET("/index.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	r.Any("/proxy", proxy.Proxy)
	r.POST("/wechat", controller.DefaultController.Post)
}
