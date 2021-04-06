/*
 * @Description: Do not edit
 * @Date: 2021-03-20 02:05:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-05 19:57:18
 */

package httpserver

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/wechat-offical-account/logger"
)

type ServerConfig struct {
	Base struct {
		Name    string `json:"name"`
		Listen  string `json:"listen"`
		RunMode string `json:"run_mode"`
	} `json:"base"`
	WeChat struct {
		Token string `json:"token"`
	} `json:"wechat"`
}

func loadConfig() (*ServerConfig, error) {
	configFile := "./conf/httpserver.toml"
	config := &ServerConfig{}
	if _, err := toml.DecodeFile(configFile, config); err != nil {
		panic(err)
	}
	return config, nil
}

func Start() {
	serverConf, _ := loadConfig()
	r := gin.Default()

	r.Use(logger.Logger())

	RouterSetup(r)
	r.Run(serverConf.Base.Listen)
}
