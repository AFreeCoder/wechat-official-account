/*
 * @Description: message 类型定义
 * @Date: 2021-03-20 19:40:20
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-18 19:25:42
 */

package message

import (
	"encoding/xml"
)

// ClientTextMsg 客户端传来的文本消息，MsgType=text
type ClientTextMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `json:"ToUserName" xml:"ToUserName"`
	FromUserName string   `json:"FromUserName" xml:"FromUserName"`
	MsgType      string   `json:"MsgType" xml:"MsgType"`
	Content      string   `json:"Content" xml:"Content"`
	MsgID        int64    `json:"MsgId" xml:"MsgId"`
	CreateTime   int64    `json:"CreateTime" xml:"CreateTime"`
}

// ClientPicMsg 客户端传来的图片消息，MsgType=image
type ClientPicMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `json:"ToUserName" xml:"ToUserName"`
	FromUserName string   `json:"FromUserName" xml:"FromUserName"`
	MsgType      string   `json:"MsgType" xml:"MsgType"`
	PicURL       string   `json:"PicUrl" xml:"PicUrl"`
	MediaID      string   `json:"MediaId" xml:"MediaId"`
	MsgID        int64    `json:"MsgId" xml:"MsgId"`
	CreateTime   int64    `json:"CreateTime" xml:"CreateTime"`
}

// ClientVoiceMsg 客户端传来的语音消息，MsgType=voice
type ClientVoiceMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `json:"ToUserName" xml:"ToUserName"`
	FromUserName string   `json:"FromUserName" xml:"FromUserName"`
	MsgType      string   `json:"MsgType" xml:"MsgType"`
	MediaID      string   `json:"MediaId" xml:"MediaId"`
	Format       string   `json:"Format" xml:"Format"`
	MsgID        int64    `json:"MsgId" xml:"MsgId"`
	CreateTime   int64    `json:"CreateTime" xml:"CreateTime"`
}

// ClientVideoMsg 客户端传来的视频/小视频消息，MsgType=video/shortvideo
type ClientVideoMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `json:"ToUserName" xml:"ToUserName"`
	FromUserName string   `json:"FromUserName" xml:"FromUserName"`
	MsgType      string   `json:"MsgType" xml:"MsgType"`
	MediaID      string   `json:"MediaId" xml:"MediaId"`
	ThumbMediaID string   `json:"ThumbMediaId" xml:"ThumbMediaId"`
	MsgID        int64    `json:"MsgId" xml:"MsgId"`
	CreateTime   int64    `json:"CreateTime" xml:"CreateTime"`
}

// ClientGeoMsg 客户端传来的地理位置消息，MsgType=location
type ClientGeoMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `json:"ToUserName" xml:"ToUserName"`
	FromUserName string   `json:"FromUserName" xml:"FromUserName"`
	MsgType      string   `json:"MsgType" xml:"MsgType"`
	LocationX    string   `json:"Location_X" xml:"Location_X"`
	LocationY    string   `json:"Location_Y" xml:"Location_Y"`
	Scale        string   `json:"Scale" xml:"Scale"`
	Lable        string   `json:"Lable" xml:"Lable"`
	MsgID        int64    `json:"MsgId" xml:"MsgId"`
	CreateTime   int64    `json:"CreateTime" xml:"CreateTime"`
}

// ClientLinkMsg 客户端传来的链接消息，MsgType=link
type ClientLinkMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `json:"ToUserName" xml:"ToUserName"`
	FromUserName string   `json:"FromUserName" xml:"FromUserName"`
	MsgType      string   `json:"MsgType" xml:"MsgType"`
	Title        string   `json:"Title" xml:"Title"`
	URL          string   `json:"Url" xml:"Url"`
	Description  string   `json:"Description" xml:"Description"`
	MsgID        int64    `json:"MsgId" xml:"MsgId"`
	CreateTime   int64    `json:"CreateTime" xml:"CreateTime"`
}

// ServerTextMsg 服务端返回的文本消息，MsgType=text
type ServerTextMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `form:"ToUserName" xml:"ToUserName"`
	FromUserName string   `form:"FromUserName" xml:"FromUserName"`
	MsgType      string   `form:"MsgType" xml:"MsgType"`
	Content      string   `form:"Content" xml:"Content"`
	CreateTime   int64    `form:"CreateTime" xml:"CreateTime"`
}

// ServerPicMsg 服务端返回的图片消息，MsgType=image
type ServerPicMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `form:"ToUserName" xml:"ToUserName"`
	FromUserName string   `form:"FromUserName" xml:"FromUserName"`
	MsgType      string   `form:"MsgType" xml:"MsgType"`
	Image        struct {
		XMLName xml.Name `xml:"Image"`
		MediaID string   `form:"MediaId" xml:"MediaId"`
	}
	CreateTime int64 `form:"CreateTime" xml:"CreateTime"`
}

// ServerVoiceMsg 服务端返回的语音消息，MsgType=voice
type ServerVoiceMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `form:"ToUserName" xml:"ToUserName"`
	FromUserName string   `form:"FromUserName" xml:"FromUserName"`
	MsgType      string   `form:"MsgType" xml:"MsgType"`
	Voice        struct {
		XMLName xml.Name `xml:"Voice"`
		MediaID string   `form:"MediaId" xml:"MediaId"`
	}
	CreateTime int64 `form:"CreateTime" xml:"CreateTime"`
}

// ServerVideoMsg 服务端返回的视频消息，MsgType=video
type ServerVideoMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `form:"ToUserName" xml:"ToUserName"`
	FromUserName string   `form:"FromUserName" xml:"FromUserName"`
	MsgType      string   `form:"MsgType" xml:"MsgType"`
	Video        struct {
		XMLName     xml.Name `xml:"Video"`
		MediaID     string   `form:"MediaId" xml:"MediaId"`
		Title       string   `json:"Title" xml:"Title"`
		Description string   `json:"Description" xml:"Description"`
	}
	CreateTime int64 `form:"CreateTime" xml:"CreateTime"`
}

// ServerMusicMsg 服务端返回的音乐消息，MsgType=music
type ServerMusicMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `form:"ToUserName" xml:"ToUserName"`
	FromUserName string   `form:"FromUserName" xml:"FromUserName"`
	MsgType      string   `form:"MsgType" xml:"MsgType"`
	Music        struct {
		XMLName      xml.Name `xml:"Music"`
		ThumbMediaID string   `form:"ThumbMediaId" xml:"ThumbMediaId"`
		Title        string   `json:"Title" xml:"Title"`
		Description  string   `json:"Description" xml:"Description"`
		MusicURL     string   `json:"MusicUrl" xml:"MusicUrl"`
		HQMusicURL   string   `json:"HQMusicUrl" xml:"HQMusicUrl"`
	}
	CreateTime int64 `form:"CreateTime" xml:"CreateTime"`
}

// ServerNewsMsg 服务端返回的图文消息，MsgType=news
type ServerNewsMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `form:"ToUserName" xml:"ToUserName"`
	FromUserName string   `form:"FromUserName" xml:"FromUserName"`
	MsgType      string   `form:"MsgType" xml:"MsgType"`
	ArticleCount uint8    `form:"ArticleCount" xml:"ArticleCount"`
	Articles     struct {
		XMLName xml.Name `xml:"Articles"`
		item    []struct {
			XMLName     xml.Name `xml:"item"`
			Title       string   `json:"Title" xml:"Title"`
			Description string   `json:"Description" xml:"Description"`
			PicURL      string   `json:"PicUrl" xml:"PicUrl"`
			URL         string   `json:"Url" xml:"Url"`
		}
	}
	CreateTime int64 `form:"CreateTime" xml:"CreateTime"`
}
