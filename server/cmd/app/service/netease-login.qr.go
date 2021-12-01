package service

import (
	"encoding/json"
	"fmt"
	"wy_music_cloud/internal/netease"
)

var options *netease.Options

type NeteaseLoginQr struct {
}

func init() {
	options = &netease.Options{
		Crypto: "weapi",
		Ua:     "pc",
	}
}

// GetUniKey 获取二维码key
func (n *NeteaseLoginQr) GetUniKey() string {
	data := make(map[string]string)
	data["type"] = "1"
	request, _ := netease.CreateRequest("POST", `https://music.163.com/weapi/login/qrcode/unikey`, data, options, nil)
	indent, err := json.MarshalIndent(request, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Print("2222" + string(indent))
	uniKey := request["unikey"].(string)

	return uniKey
}

// CheckQr 检测二维码状态,成功返回 true,cookie 800 二维码不存在或已过期 801等待扫码 802 等待确认 803 登录成功
func (n *NeteaseLoginQr) CheckQr(uniKey string) (bool, string) {
	data := make(map[string]string)
	data["type"] = "1"
	data["key"] = uniKey
	request, cookies := netease.CreateRequest("POST", `https://music.163.com/weapi/login/qrcode/client/login`, data, options, nil)
	if request["code"].(float64) != 803 {
		return false, request["message"].(string)
	}
	cookiesStr := ""
	for _, cookie := range cookies {
		if cookiesStr != "" {
			cookiesStr = cookiesStr + ";"
		}
		cookiesStr = cookiesStr + cookie.String()
	}
	return true, cookiesStr
}
