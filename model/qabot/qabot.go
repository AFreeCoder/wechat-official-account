/*
 * @Description: Do not edit
 * @Date: 2021-03-28 02:20:18
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-19 01:53:57
 */

package qabot

import "github.com/gin-gonic/gin"

// QA 问答，对参数text做出回答，答案暂时只有文本类型
func QA(c *gin.Context, text string) (interface{}, error) {
	switch text {
	default:
		return "自动问答功能开发中，请耐心等待~", nil
	}
}
