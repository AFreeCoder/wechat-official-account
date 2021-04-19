/*
 * @Description: Do not edit
 * @Date: 2021-04-19 02:05:51
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-19 02:34:03
 */

package global

import "github.com/BurntSushi/toml"

type wechatConfig struct {
	Token string `json:"token"`
}

// WeChatConf 全局配置
var WeChatConf = wechatConfig{}

func init() {
	configFile := "./conf/wechat.toml"
	if _, err := toml.DecodeFile(configFile, &WeChatConf); err != nil {
		panic(err)
	}
}
