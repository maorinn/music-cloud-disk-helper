package netease

import (
	"testing"
	"wy_music_cloud/config"
)

func TestUploadImg(t *testing.T) {
	loginCellphoneService := LoginCellphoneService{
		Phone:    config.Config.MusicAccount.Phone,
		Password: config.Config.MusicAccount.Password,
	}
	_, cookies := loginCellphoneService.LoginCellphone()
	UploadImg(cookies, config.Config.HomePath+"/tmp/img.png")
}
