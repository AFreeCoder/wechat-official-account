/*
 * @Description: Do not edit
 * @Date: 2021-03-20 13:27:25
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-04-19 02:34:31
 */

package util

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/wechat-official-account/global"
)

// VerifyParams 签名校验
func VerifyParams(signature, timestamp, nonce string) (bool, error) {
	var array = []string{global.WeChatConf.Token, timestamp, nonce}
	sort.Strings(array)
	tmpStr := strings.Join(array, "")
	h := sha1.New()
	if _, err := h.Write([]byte(tmpStr)); err != nil {
		return false, err
	}
	sign := hex.EncodeToString(h.Sum(nil))
	return signature == sign, nil
}
