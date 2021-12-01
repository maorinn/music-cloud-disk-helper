package main

import (
	"testing"
	"wy_music_cloud/cmd/app/service"
	"wy_music_cloud/internal/bilibili"
)

var c *bilibili.BiliClient

func init() {
	c, _ = bilibili.NewBiliClient(&bilibili.BiliSetting{
		DebugMode: true,
	})
}

func TestBiliClient_VideoDescription(t *testing.T) {

	description, err := c.VideoGetDescription(759949922)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("desc: %s", description)
}

func TestGetLoginUrl(t *testing.T) {
	s := service.BiliLoginQr{BiliClient: c}
	url, err := s.GetLoginUrl()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("url: %s", url)
}

func TestGetUniKey(t *testing.T) {
	qrService := service.NeteaseLoginQr{}
	key := qrService.GetUniKey()
	t.Logf("key: %s", key)
}
