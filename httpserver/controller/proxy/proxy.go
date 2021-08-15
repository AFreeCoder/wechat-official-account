/*
 * @Description: Do not edit
 * @Date: 2021-08-12 15:26:08
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-08-15 20:06:00
 */

package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	host        = "https://www.jisilu.cn"
	upstreamURL = "X-Upstream-Url"
)

// Proxy 验证消息的确来自微信服务器
func Proxy(c *gin.Context) {
	url := strings.TrimLeft(c.Request.Header.Get(upstreamURL), "?") + "?" + c.Request.URL.RawQuery
	c.Request.Header.Del(upstreamURL)
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	req.Header = c.Request.Header.Clone()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	// header 复制
	for k := range resp.Header.Clone() {
		v := resp.Header.Get(k)
		c.Writer.Header().Set(k, v)
	}
	c.Writer.WriteHeader(resp.StatusCode)
	c.Writer.Write(body)
	return
}
