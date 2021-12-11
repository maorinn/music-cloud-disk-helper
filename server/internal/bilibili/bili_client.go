package bilibili

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"wy_music_cloud/common"
	"wy_music_cloud/utils"
)

type BiliClient struct {
	Me   *Account
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

// NewBiliClient 带有账户Cookie的Client，用于访问私人操作API
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

	//if bili.auth != nil {
	//	fmt.Println("1111111111111" + bili.auth.BiliJCT)
	//	account, err := bili.GetMe()
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	bili.Me = account
	//}

	return bili, nil
}

// GetLoginUrl 申请二维码URL及扫码密钥
func (b *BiliClient) GetLoginUrl() (*QrLoinUrl, error) {
	resp, err := b.RawParse(common.BiliPassportURL, "qrcode/getLoginUrl", "GET", nil)
	if err != nil {
		return nil, err
	}
	var qrLoinUrl *QrLoinUrl
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
func (b *BiliClient) GetLoginInfo(qrLoinUrl *QrLoinUrl) (*common.Response, error) {
	resp, err := b.RawParse(common.BiliPassportURL, "qrcode/getLoginInfo", "POST", map[string]string{
		"oauthKey": qrLoinUrl.OauthKey,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetMe 获取个人基本信息
func (b *BiliClient) GetMe() (*Account, error) {
	resp, err := b.RawParse(BiliApiURL,
		"x/member/web/account",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}
	var account *Account
	if err = json.Unmarshal(resp.Data, &account); err != nil {
		return nil, err
	}
	return account, nil
}

// VideoGetPageList 获取分P列表
func (b *BiliClient) VideoGetPageList(bvid string) ([]*VideoPage, error) {
	resp, err := b.RawParse(BiliApiURL,
		"x/player/pagelist",
		"GET",
		map[string]string{
			"bvid": bvid,
		},
	)
	if err != nil {
		return nil, err
	}
	var list []*VideoPage
	if err = json.Unmarshal(resp.Data, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// VideoGetPlayURL 获取视频取流地址
// 所有参数、返回信息和取流方法的说明请直接前往：https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/videostream_url.md
func (b *BiliClient) VideoGetPlayURL(bvid string, cid int64, qn int, fnval int) (*VideoPlayURLResult, error) {
	resp, err := b.RawParse(
		BiliApiURL,
		"x/player/playurl",
		"GET",
		map[string]string{
			"bvid":  bvid,
			"cid":   strconv.FormatInt(cid, 10),
			"qn":    strconv.Itoa(qn),
			"fnval": strconv.Itoa(fnval),
			"fnver": "0",
			"fourk": "1",
		},
	)
	if err != nil {
		return nil, err
	}
	var r *VideoPlayURLResult
	if err = json.Unmarshal(resp.Data, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// Down 下载资源
func (b *BiliClient) Down(url, path string, fileName string) error {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("range", "bytes=0-")
	request.Header.Add("referer", "https://www.bilibili.com")
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	exists := utils.Exists(path)
	if exists == false {
		os.Mkdir(path, os.ModePerm)
	}
	f, err := os.Create(path + "/" + fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	// mao 或许可以直接上传
	if _, err := io.Copy(f, res.Body); err != nil {
		return err
	}
	return nil
}

// SetClient 设置Client,可以用来更换代理等操作
func (b *BiliClient) SetClient(client *http.Client) {
	b.client = client
}

// SetUA 设置UA
func (b *BiliClient) SetUA(ua string) {
	b.ua = ua
}

// Raw base末尾带/
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

// RawParse base末尾带/
func (b *BiliClient) RawParse(base, endpoint, method string, payload map[string]string) (*common.Response, error) {
	raw, err := b.Raw(base, endpoint, method, payload)
	if err != nil {
		return nil, err
	}
	return b.parse(raw)
}

// GetCookieAuth 获取Cookie信息
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

// VideoGetView 获取视频基础信息
func (b *BiliClient) VideoGetView(bvid string) (*VideoView, error) {
	resp, err := b.RawParse(common.BiliApiURL,
		"x/web-interface/view",
		"GET",
		map[string]string{
			"bvid": bvid,
		},
	)
	if err != nil {
		return nil, err
	}
	var videoView *VideoView
	if err = json.Unmarshal(resp.Data, &videoView); err != nil {
		return nil, err
	}
	return videoView, nil
}
