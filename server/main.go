package main

import (
	"encoding/json"
	"fmt"
	"wy_music_cloud/config"
	_ "wy_music_cloud/config"
	"wy_music_cloud/plugins"
	"wy_music_cloud/service"
)

func main() {
	loginCellphoneService := service.LoginCellphoneService{
		Phone:    config.Config.MusicAccount.Phone,
		Password: config.Config.MusicAccount.Password,
	}
	cellphone, cookies := loginCellphoneService.LoginCellphone()
	config.Config.MusicAccount.Cookie = cookies
	b, err := json.MarshalIndent(cellphone, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// 上传
	cloud := plugins.UploadCloud{
		FilePath: "./夜曲.mp3",
		Cookies:  config.Config.MusicAccount.Cookie,
	}
	cloud.UploadCloud()
}
