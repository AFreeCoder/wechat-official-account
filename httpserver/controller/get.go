/*
 * @Description: Do not edit
 * @Date: 2021-03-20 12:41:39
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-03-21 01:23:34
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wechat-offical-account/library/util"
)

// GetParams 请求参数
type GetParams struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	Echostr   string `form:"echostr"`
}

// Message 公众号消息
type Message struct {
	ToUserName   string `form:"ToUserName"`
	FromUserName string `form:"FromUserName"`
	MsgType      string `form:"MsgType"`
	Content      string `form:"Content"`
	MsgID        string `form:"MsgId"`
	CreateTime   string `form:"CreateTime"`
}

// Get 验证消息的确来自微信服务器
func (bc *BaseController) Get(c *gin.Context) {
	var params = &GetParams{}
	c.ShouldBind(params)
	verfifyRes, err := util.VerifyParams(params.Signature, params.Timestamp, params.Nonce)
	if err != nil || verfifyRes != true {
		bc.ResponseString(c, "hhhhh")
	}

	bc.ResponseString(c, params.Echostr)
	return
}
