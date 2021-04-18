/*
 * @Description: 关键词回复
 * @Date: 2021-04-18 17:52:38
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-19 01:06:51
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
		return ServerTextMsg{
			ToUserName:   from,
			FromUserName: to,
			MsgType:      "text",
			Content:      "可以成功。因为这个时候 readerCount 已经被恢复成正整数。RLock 操作不会有任何影响。",
			CreateTime:   time.Now().Unix(),
		}, true
	default:
		return ServerTextMsg{}, false
	}
}

func eventReply(event string, from, to string) (interface{}, bool) {
	switch event {
	case "subscribe":
		return ServerTextMsg{
			ToUserName:   from,
			FromUserName: to,
			MsgType:      "text",
			Content: `欢迎关注【码农的自由之路】！

			这里有投资理财技能的分享： <a href="https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg5NjU0MzU1Mg==&action=getalbum&album_id=1682265629497589761#wechat_redirect">理财，咋入门</a>；也有编程相关知识的分享： <a href="https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg5NjU0MzU1Mg==&action=getalbum&album_id=1682264409559105537#wechat_redirect">Go Go Go</a>；！
			
			一个996程序员的财务自由日志： <a href="https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg5NjU0MzU1Mg==&action=getalbum&album_id=1660571680026492929#wechat_redirect">财务自由实证</a> .
			
			左手代码，右手理财，一起见证自由之路！
			
			感谢关注，让我们一路同行~`,
			CreateTime: time.Now().Unix(),
		}, true
	default:
		return ServerTextMsg{
			ToUserName:   from,
			FromUserName: to,
			MsgType:      "text",
			Content:      "success",
			CreateTime:   time.Now().Unix(),
		}, false
	}
}
