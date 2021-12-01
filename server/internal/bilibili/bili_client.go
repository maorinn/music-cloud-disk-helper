package bilibili

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"wy_music_cloud/common"
)

type BiliClient struct {
	Me   *common.Account
	auth *CookieAuth

	*baseClient
}

type CookieAuth struct {
	DedeUserID      string // DedeUserID
	DedeUserIDCkMd5 string // DedeUserID__ckMd5
	SESSDATA        string // SESSDATA
	BiliJCT         string // bili_jct
}

type BiliSetting struct {
	// Cookie
	Auth *CookieAuth
	// 自定义http client
	//
	// 默认为 http.http.DefaultClient
	Client *http.Client
	// Debug模式 true将输出请求信息 false不输出
	//
	// 默认false
	DebugMode bool
	// 自定义UserAgent
	//
	// 默认Chrome随机Agent
	UserAgent string
}

// NewBiliClient
//
// 带有账户Cookie的Client，用于访问私人操作API
func NewBiliClient(setting *BiliSetting) (*BiliClient, error) {
	//if setting.Auth == nil {
	//	return nil, errors.New("auth cannot be nil")
	//}
	bili := &BiliClient{
		auth: setting.Auth,
		baseClient: newBaseClient(&baseSetting{
			Client:    setting.Client,
			DebugMode: setting.DebugMode,
			UserAgent: setting.UserAgent,
			Prefix:    "BiliClient ",
		}),
	}

	if bili.auth != nil {
		account, err := bili.GetMe()
		if err != nil {
			return nil, err
		}

		bili.Me = account
	}

	return bili, nil
}

// GetMe
//
// 获取个人基本信息
func (b *BiliClient) GetMe() (*common.Account, error) {
	resp, err := b.RawParse(common.BiliApiURL,
		"x/member/web/account",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}
	var account *common.Account
	if err = json.Unmarshal(resp.Data, &account); err != nil {
		return nil, err
	}
	return account, nil
}

// SetClient
//
// 设置Client,可以用来更换代理等操作
func (b *BiliClient) SetClient(client *http.Client) {
	b.client = client
}

// SetUA
//
// 设置UA
func (b *BiliClient) SetUA(ua string) {
	b.ua = ua
}

// Raw
//
// base末尾带/
func (b *BiliClient) Raw(base, endpoint, method string, payload map[string]string) ([]byte, error) {
	raw, err := b.raw(base, endpoint, method, payload,
		func(d *url.Values) {
			switch method {
			case "POST":
				if b.auth != nil {
					d.Add("csrf", b.auth.BiliJCT)
				}

			}
		},
		func(r *http.Request) {
			if b.auth != nil {
				r.Header.Add("Cookie", fmt.Sprintf("DedeUserID=%s;SESSDATA=%s;DedeUserID__ckMd5=%s",
					b.auth.DedeUserID, b.auth.SESSDATA, b.auth.DedeUserIDCkMd5))
			}

		})
	if err != nil {
		return nil, err
	}
	return raw, nil
}

// RawParse
//
// base末尾带/
func (b *BiliClient) RawParse(base, endpoint, method string, payload map[string]string) (*common.Response, error) {
	raw, err := b.Raw(base, endpoint, method, payload)
	if err != nil {
		return nil, err
	}
	return b.parse(raw)
}

// GetCookieAuth
// 获取Cookie信息
func (b *BiliClient) GetCookieAuth() *CookieAuth {
	return b.auth
}

// GetNavInfo
// 获取我的导航栏信息(大部分的用户信息都在这里了)
func (b *BiliClient) GetNavInfo() (*common.NavInfo, error) {
	resp, err := b.RawParse(
		common.BiliApiURL,
		"x/web-interface/nav",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}
	var info *common.NavInfo
	if err = json.Unmarshal(resp.Data, &info); err != nil {
		return nil, err
	}
	return info, nil
}

// VideoGetDescription 获取稿件简介
func (b *BiliClient) VideoGetDescription(aid int64) (string, error) {
	resp, err := b.RawParse(common.BiliApiURL,
		"x/web-interface/archive/desc",
		"GET",
		map[string]string{
			"aid": strconv.FormatInt(aid, 10),
		},
	)
	if err != nil {
		return "", err
	}
	var desc string
	if err = json.Unmarshal(resp.Data, &desc); err != nil {
		return "", err
	}
	return desc, nil
}
