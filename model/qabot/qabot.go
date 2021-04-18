/*
 * @Description: Do not edit
 * @Date: 2021-03-28 02:20:18
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-18 19:29:22
 */

package qabot

import "github.com/gin-gonic/gin"

// QA 问答，对参数text做出回答，答案暂时只有文本类型
func QA(c *gin.Context, text string) (interface{}, error) {
	switch text {
	default:
		return "小A暂时无法回答这个问题哦", nil
	}
}
