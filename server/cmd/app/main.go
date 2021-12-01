package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"wy_music_cloud/cmd/app/handlers"
	_ "wy_music_cloud/config"
)

func main() {
	gin.SetMode("debug")
	routersInit := handlers.InitRouter()
	routersInit.Run(":8000")
	log.Printf("[info] start http server listening %s", ":8000")
}

//func main() {
//	loginCellphoneService := service.LoginCellphoneService{
//		Phone:    config.Config.MusicAccount.Phone,
//		Password: config.Config.MusicAccount.Password,
//	}
//	cellphone, cookies := loginCellphoneService.LoginCellphone()
//	config.Config.MusicAccount.Cookie = cookies
//	b, err := json.MarshalIndent(cellphone, "", " ")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(b))
//	// 上传
//	cloud := plugins.UploadCloud{
//		FilePath: "./夜曲.mp3",
//		Cookies:  config.Config.MusicAccount.Cookie,
//	}
//	cloud.UploadCloud()
//}
