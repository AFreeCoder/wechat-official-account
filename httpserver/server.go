/*
 * @Description: Do not edit
 * @Date: 2021-03-20 02:05:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-08-12 15:19:41
 */

package httpserver

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/wechat-official-account/logger"
)

// ServerConfig server 配置
type ServerConfig struct {
	Name    string `json:"name"`
	Listen  string `json:"listen"`
	RunMode string `json:"run_mode"`
}

func loadConfig() (*ServerConfig, error) {
	configFile := "./conf/httpserver.toml"
	config := &ServerConfig{}
	if _, err := toml.DecodeFile(configFile, config); err != nil {
		panic(err)
	}
	return config, nil
}

// Start 入口
func Start() {
	serverConf, _ := loadConfig()
	r := gin.Default()
	r.Use(logger.Logger())
	r.Static("/static/", "./static/")
	r.LoadHTMLFiles("./static/index.html")
	RouterSetup(r)
	r.Run(serverConf.Listen)
}
