package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wy_music_cloud/cmd/app/handlers/base"
	"wy_music_cloud/config"
	"wy_music_cloud/internal/bilibili"
	"wy_music_cloud/internal/netease"
	"wy_music_cloud/utils"
)

var path = config.Config.HomePath + "/tmp"

type UploadSongDto struct {
	Bvid string `json:"bvid"`
}

func UploadSong(c *gin.Context) {
	dto := &UploadSongDto{}
	c.BindJSON(&dto)
	// 获取cookie信息
	s := &base.SansAuth{
		NeteaseCookieStr: c.GetHeader("netease_cookie_str"),
		BiliCookieStr:    c.GetHeader("bili_cookie_str"),
	}
	// 构造网易云cookies对象
	neteaseCookies := c.Request.Cookies()
	split := strings.Split(s.NeteaseCookieStr, ";")
	for _, str := range split {
		// 分离kv
		i := strings.Split(str, "=")
		fmt.Print("33333:" + i[0] + "vvvvvv" + i[1])
		neteaseCookies = append(neteaseCookies, &http.Cookie{Name: i[0], Value: i[1]})
	}
	// 构造Bili验证对象
	biliAuth := &bilibili.CookieAuth{
		DedeUserID:      utils.GetBetweenStr(s.BiliCookieStr, "DedeUserID=", "&"),
		SESSDATA:        utils.GetBetweenStr(s.BiliCookieStr, "SESSDATA=", "&"),
		DedeUserIDCkMd5: utils.GetBetweenStr(s.BiliCookieStr, "DedeUserID__ckMd5=", "&"),
		BiliJCT:         utils.GetBetweenStr(s.BiliCookieStr, "bili_jct=", "&"),
	}
	// 构造Bili client
	b, _ := bilibili.NewBiliClient(&bilibili.BiliSetting{
		DebugMode: true,
		Auth:      biliAuth,
	})
	// 构造网易云client
	neteaseOptions := &netease.Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: neteaseCookies,
	}
	var _ = neteaseOptions
	// 分P列表
	pList, err := b.VideoGetPageList(dto.Bvid)
	if err != nil {
		base.Re(c, -1, err.Error(), nil)
	}
	// 获取音频下载地址
	playURL, err := b.VideoGetPlayURL(dto.Bvid, pList[0].CID, 80, 16)
	if err != nil {
		base.Re(c, -1, err.Error(), nil)
	}
	biliAudioDownloadUrl := playURL.Dash.Audio[0].BaseURL
	// 获取视频基础信息
	view, err := b.VideoGetView(dto.Bvid)
	view.Title = strings.Replace(view.Title,"/","-",-1)
	// 生成文件名
	fileName := view.Title + ".mp3"
	// 判断文件是否存在
	exists := utils.Exists(path + "/" + fileName)
	if !exists {
		// 下载
		err = b.Down(biliAudioDownloadUrl, path, fileName)
		if err != nil {
			base.Re(c, -1, err.Error(), nil)
		}
	}
	// 上传网易云音乐
	cloud := netease.UploadCloud{
		FilePath: path + "/" + fileName,
		Cookies:  neteaseCookies,
		SongMetadata: &netease.SongMetadata{
			Title:  view.Title,
			Artist: view.Owner.Name,
			Album:  "未知专辑",
		},
	}
	cloud.UploadCloud()
	base.Re(c, 0, "success", map[string]string{
		"fileName": fileName,
	})
}
