/*
 * @Description: Do not edit
 * @Date: 2021-03-20 02:05:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-19 02:30:31
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

	RouterSetup(r)
	r.Run(serverConf.Listen)
}
