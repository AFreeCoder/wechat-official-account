/*
 * @Description: 消息处理
 * @Date: 2021-03-27 17:11:48
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-19 01:07:20
 */

package message

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

// EventHandler 事件消息处理
func EventHandler(c *gin.Context, event string) (interface{}, error) {
	clientMsg := &ClientTextMsg{}
	if err := c.ShouldBindBodyWith(clientMsg, binding.XML); err != nil {
		log.Warn("bind body err: ", err)
		return nil, err
	}

	// 先匹配固定回复
	resp, ok := eventReply(event, clientMsg.FromUserName, clientMsg.ToUserName)
	if ok {
		return resp, nil
	}
	// TODO: 对其它事件进行处理
	return resp, fmt.Errorf("no reply")
}

// MsgHandler 被动回复用户消息
func MsgHandler(c *gin.Context, msgType string) (interface{}, error) {
	switch msgType {
	case "text":
		return TextHandler(c)
	}
	return nil, nil
}

// TextHandler 文本消息处理
func TextHandler(c *gin.Context) (interface{}, error) {
	clientMsg := &ClientTextMsg{}
	if err := c.ShouldBindBodyWith(clientMsg, binding.XML); err != nil {
		log.Warn("bind body err: ", err)
		return nil, err
	}

	// 先看配置关键词回复
	if resp, ok := keyReplay(clientMsg.Content, clientMsg.FromUserName, clientMsg.ToUserName); ok {
		return resp, nil
	}

	// 暂时取消兜底回复
	return "success", fmt.Errorf("no replay")
}
