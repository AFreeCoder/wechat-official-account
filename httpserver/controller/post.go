/*
 * @Description: Do not edit
 * @Date: 2021-03-20 19:28:21
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-18 19:39:30
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/wechat-official-account/httpserver/controller/message"
	"github.com/wechat-official-account/library/util"
)

// MessageType 获取消息类型，用于分流，让不同的handler处理
type MessageType struct {
	Event   string `form:"Event"`
	MsgType string `form:"MsgType"`
}

// Post 处理微信的post请求
func (bc *BaseController) Post(c *gin.Context) {
	var params = &GetParams{}
	c.ShouldBindQuery(params)
	verfifyRes, err := util.VerifyParams(params.Signature, params.Timestamp, params.Nonce)
	if err != nil || !verfifyRes {
		bc.ResponseString(c, "success")
	}

	var m = &MessageType{}
	c.ShouldBindBodyWith(m, binding.XML)

	var resp interface{}

	if m.Event != "" {
		resp, err = message.EventHandler(c, m.Event)
	} else {
		resp, err = message.MsgHandler(c, m.MsgType)
	}
	if err == nil {
		bc.ResponseXML(c, resp)
	} else {
		bc.ResponseString(c, resp.(string))
	}
}
