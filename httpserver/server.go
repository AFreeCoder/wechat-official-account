/*
 * @Description: Do not edit
 * @Date: 2021-03-20 02:05:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-03-28 02:13:42
 */

package httpserver

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()

	RouterSetup(r)

	r.Run(":80")
}
