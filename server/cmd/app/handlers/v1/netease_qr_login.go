package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wy_music_cloud/cmd/app/handlers/base"
	"wy_music_cloud/cmd/app/service"
)

// GetUniKey 获取登录二维码key
func GetUniKey(c *gin.Context) {
	qr := service.NeteaseLoginQr{}
	key := qr.GetUniKey()
	base.Re(c, 0, "success", &map[string]string{
		"key": key,
	})
}

// GetCheckQr 获取二维码扫描状态
func GetCheckQr(c *gin.Context) {
	key := c.Query("key")
	fmt.Print("222" + key)
	qr := service.NeteaseLoginQr{}
	status, cookie := qr.CheckQr(key)
	base.Re(c, 0, "success", &map[string]interface {
	}{
		"status": status,
		"cookie": cookie,
	})
}
