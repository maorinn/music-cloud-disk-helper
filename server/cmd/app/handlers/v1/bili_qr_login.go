package v1

import (
	"github.com/gin-gonic/gin"
	"wy_music_cloud/cmd/app/handlers/base"
	"wy_music_cloud/internal/bilibili"
)

var b *bilibili.BiliClient

func init() {
	b, _ = bilibili.NewBiliClient(&bilibili.BiliSetting{
		DebugMode: true,
	})
}

// GetLoginUrl 申请二维码信息
func GetLoginUrl(c *gin.Context) {
	r, err := b.GetLoginUrl()
	if err != nil {
		base.Re(c, -1, err.Error(), nil)
	}
	base.Re(c, 0, "success", r)
}

// GetLoginInfo 获取扫码结果
func GetLoginInfo(c *gin.Context) {
	var q *bilibili.QrLoinUrl
	c.BindJSON(&q)
	r, err := b.GetLoginInfo(q)
	if err != nil {
		base.Re(c, -1, err.Error(), nil)
	}
	base.Re(c, 0, "success", r)
}
