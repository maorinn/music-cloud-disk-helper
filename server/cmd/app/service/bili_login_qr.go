package service

import (
	"encoding/json"
	"wy_music_cloud/common"
	"wy_music_cloud/internal/bilibili"
)

type BiliLoginQr struct {
	BiliClient *bilibili.BiliClient
}

// GetLoginUrl 申请二维码URL及扫码密钥
func (b *BiliLoginQr) GetLoginUrl() (*common.QrLoinUrl, error) {
	resp, err := b.BiliClient.RawParse(common.BiliPassportURL, "qrcode/getLoginUrl", "GET", nil)
	if err != nil {
		return nil, err
	}
	var qrLoinUrl *common.QrLoinUrl
	if err = json.Unmarshal(resp.Data, &qrLoinUrl); err != nil {
		return nil, err
	}

	return qrLoinUrl, nil
}

// GetLoginInfo 获取登录结果
//status	bool	扫码是否成功	true：成功
//false：未成功
//data	正确时：obj
//错误时：num	正确时：游戏分站url
//错误时：错误代码	未成功时：
//-1：密钥错误
//-2：密钥超时
//-4：未扫描a
//-5：未确认
func (b *BiliLoginQr) GetLoginInfo(qrLoinUrl *common.QrLoinUrl) (*common.Response, error) {
	resp, err := b.BiliClient.RawParse(common.BiliPassportURL, "qrcode/getLoginInfo", "POST", map[string]string{
		"oauthKey": qrLoinUrl.OauthKey,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
