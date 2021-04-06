/*
 * @Description: 消息处理
 * @Date: 2021-03-27 17:11:48
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-05 01:35:24
 */
package message

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/wechat-offical-account/model/text"
)

// TextHandle 文本消息处理
func TextHandle(c *gin.Context) (interface{}, error) {
	clientMsg := &ClientTextMsg{}
	if err := c.ShouldBindBodyWith(clientMsg, binding.XML); err != nil {
		// Todo: log
		return nil, err
	}
	fmt.Println(clientMsg)
	serverMsg := &ServerTextMsg{}
	serverMsg.ToUserName = clientMsg.FromUserName
	serverMsg.FromUserName = clientMsg.ToUserName
	serverMsg.MsgType = "text"
	serverMsg.CreateTime = time.Now().Unix()
	content, err := text.QA(c, clientMsg.Content)
	if err != nil {
		// Todo: log
	}
	serverMsg.Content = content.(string)
	return serverMsg, nil
}
