package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wy_music_cloud/cmd/app/handlers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(Cors())
	r.Use(gin.Recovery())
	r.GET("/ping", pong)
	apiv1 := r.Group("/api/v1")
	apiv1.POST("/UploadSong",v1.UploadSong)
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
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}