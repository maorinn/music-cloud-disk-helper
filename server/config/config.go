package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)
type MusicAccountConf struct {
	Phone    string `mapstructure:"phone"`
	Password string `mapstructure:"password"`
	Cookie []*http.Cookie
}

type BiliAccountConf struct {
	Cookie string `yaml:"cookie"`
}
type Conf struct {
	MusicAccount MusicAccountConf `mapstructure:"music_account"`
	BiliAccount BiliAccountConf `mapstructure:"bili_account"`
}


var Config = new(Conf)
func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := viper.Unmarshal(&Config); nil != err {
		log.Fatalf("赋值配置对象失败，异常信息：%v", err)
	}
	fmt.Printf("config: %#v\n", Config)
}
