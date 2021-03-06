package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"wy_music_cloud/internal/netease"
)

type LoginCellphoneService struct {
	Phone       string `json:"phone" form:"phone"`
	Countrycode string `json:"countrycode" form:"countrycode"`
	Password    string `json:"password" form:"password"`
	Md5password string `json:"md5_password" form:"md5_password"`
}

func (service *LoginCellphoneService) LoginCellphone() (map[string]interface{}, []*http.Cookie) {

	// 获得所有cookie
	cookies := []*http.Cookie{{Name: "os", Value: "pc"},{Name: "appver",Value: "2.7.1.198277"}}
	//cookiesOS :=
	//cookies = append(cookies, cookiesOS)

	options := &netease.Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: cookies,
	}
	data := make(map[string]string)

	data["phone"] = service.Phone
	if service.Countrycode != "" {
		data["countrycode"] = service.Countrycode
	}
	if service.Password != "" {
		h := md5.New()
		h.Write([]byte(service.Password))
		data["password"] = hex.EncodeToString(h.Sum(nil))
	} else {
		data["password"] = service.Md5password
	}
	data["rememberLogin"] = "true"

	//reBody, cookies := util.CreateRequest("POST", `https://www.httpbin.org/post`, data, options)
	reBody, cookies := netease.CreateRequest("POST", `https://music.163.com/weapi/login/cellphone`, data, options, nil)

	cookiesStr := ""

	for _, cookie := range cookies {
		if cookiesStr != "" {
			cookiesStr = cookiesStr + ";"
		}
		cookiesStr = cookiesStr + cookie.String()
		//c.SetCookie(cookie.Name, cookie.Value, 60*60*24, "", cookie.Domain, false, false)
	}

	reBody["cookie"] = cookiesStr
	fmt.Print("888888:"+cookiesStr)
	return reBody, cookies
}
