package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wy_music_cloud/cmd/app/handlers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", pong)
	apiBiliv1 := r.Group("/api/v1/bili")
	{
		//获取二维码url和key
		apiBiliv1.GET("/LoginUrl", v1.GetLoginUrl)
		//获取二维码扫描结果
		apiBiliv1.POST("/LoginInfo", v1.GetLoginInfo)
	}
	apiNeteasev1 := r.Group("/api/v1/netease")
	{
		apiNeteasev1.GET("qrKey", v1.GetUniKey)
		apiNeteasev1.GET("checkQr", v1.GetCheckQr)
	}
	return r
}
func pong(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}
