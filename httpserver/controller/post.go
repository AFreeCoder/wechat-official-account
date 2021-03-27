/*
 * @Description: Do not edit
 * @Date: 2021-03-20 19:28:21
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-03-28 02:37:03
 */
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/wechat-offical-account/httpserver/controller/message"
	"github.com/wechat-offical-account/library/util"
)

type MessageType struct {
	MsgType string `form:"MsgType"`
}

// Post 处理微信的post请求
func (bc *BaseController) Post(c *gin.Context) {
	var params = &GetParams{}
	c.ShouldBindQuery(params)
	verfifyRes, err := util.VerifyParams(params.Signature, params.Timestamp, params.Nonce)
	if err != nil || verfifyRes != true {
		bc.ResponseString(c, "hhhhh")
	}

	var m = &MessageType{}
	c.ShouldBindBodyWith(m, binding.XML)

	var resp interface{}

	switch m.MsgType {
	case "text":
		resp, err = message.TextHandle(c)
	default:
		resp, err = message.TextHandle(c)
	}
	bc.ResponseXML(c, resp)
}
