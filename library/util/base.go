/*
 * @Description: Do not edit
 * @Date: 2021-03-20 13:27:25
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2021-03-21 01:43:29
 */

package util

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

// VerifyParams 签名校验
func VerifyParams(signature, timestamp, nonce string) (bool, error) {
	token := "123456"
	var array = []string{token, timestamp, nonce}
	sort.Strings(array)
	tmpStr := strings.Join(array, "")
	h := sha1.New()
	if _, err := h.Write([]byte(tmpStr)); err != nil {
		return false, err
	}
	sign := hex.EncodeToString(h.Sum(nil))
	return signature == sign, nil
}
