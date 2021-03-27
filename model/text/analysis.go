/*
 * @Description: Do not edit
 * @Date: 2021-03-28 02:20:18
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-03-28 02:26:31
 */

package text

import "github.com/gin-gonic/gin"

// QA 问答，对参数text做出回答，答案暂时只有文本类型
func QA(c *gin.Context, text string) (interface{}, error) {
	switch text {
	case "哈哈":
		return "无聊", nil
	default:
		return "小A暂时无法理解您的问题，抱歉", nil
	}
}
