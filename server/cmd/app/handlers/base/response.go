package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SansAuth 外部 鉴权信息
type SansAuth struct {
	NeteaseCookieStr string `json:"netease_cookie_str"`
	BiliCookieStr string `json:"bili_cookie_str"`
}

func Re(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
