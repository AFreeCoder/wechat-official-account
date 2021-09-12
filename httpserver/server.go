/*
 * @Description: Do not edit
 * @Date: 2021-03-20 02:05:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-09-13 00:16:48
 */

package httpserver

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/wechat-official-account/logger"
)

// ServerConfig server 配置
type ServerConfig struct {
	Name           string `toml:"name"`
	Listen         string `toml:"listen"`
	PemPath        string `toml:"pem_path"`
	PrivateKeyPath string `toml:"private_key_path"`
	RunMode        string `toml:"run_mode"`
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
	r.RunTLS(serverConf.Listen, serverConf.PemPath, serverConf.PrivateKeyPath)
}
