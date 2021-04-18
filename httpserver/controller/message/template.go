/*
 * @Description: 关键词回复
 * @Date: 2021-04-18 17:52:38
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-18 19:22:55
 */

package message

import "time"

// keyReplay 关键词回复
func keyReplay(key string, from, to string) (interface{}, bool) {
	switch key {
	case "搞事情":
		return ServerTextMsg{
			ToUserName:   from,
			FromUserName: to,
			MsgType:      "text",
			Content:      "Hi 你好，我打算开发一些简单有效的投资策略，希望能通过实践加快投资体系的建立。这个过程中，我在寻找一些志同道合的小伙伴，一起合作，共同进步。如果你也正有此意，欢迎加我微信：AFreeCoder01。\n让我们一起来搞事情！",
			CreateTime:   time.Now().Unix(),
		}, true
	case "Unlock答案":
		return ServerTextMsg{}, true
	default:
		return ServerTextMsg{}, false
	}
}
