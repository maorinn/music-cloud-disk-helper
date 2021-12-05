package netease

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"hash"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type SongMetadata struct {
	Title  string // 歌曲名称
	Artist string // 艺术家
	Album  string // 专辑
}

type TokenRes struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Result  Result      `json:"result"`
}
type Result struct {
	Bucket     string `json:"bucket"`
	DocID      string `json:"docId"`
	ObjectKey  string `json:"objectKey"`
	OuterURL   string `json:"outerUrl"`
	ResourceID int64  `json:"resourceId"`
	Token      string `json:"token"`
}

type UploadCloud struct {
	FilePath     string
	Cookies      []*http.Cookie
	md5h         hash.Hash
	SongMetadata *SongMetadata
}

var (
	songName = ""
	artist   = "未知艺术家"
	album    = "未知专辑"
	stat     os.FileInfo
)

func (service *UploadCloud) UploadCloud() map[string]interface{} {
	service.Cookies = append(service.Cookies, &http.Cookie{Name: "os", Value: "pc"}, &http.Cookie{Name: "appver", Value: "2.7.1.198277"})
	file, err := os.Open(service.FilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	service.md5h = md5.New()
	io.Copy(service.md5h, file)
	if err != nil {
		panic(err)
	}
	stat, err = os.Stat(service.FilePath)

	if err != nil {
		panic(err)
	}
	songName = strings.Split(stat.Name(), ".mp3")[0]
	fmt.Print(stat.Size())
	data := make(map[string]interface{}, 0)
	data["bitrate"] = "999000"
	data["ext"] = ""
	data["length"] = strconv.FormatInt(stat.Size(), 10)
	data["md5"] = hex.EncodeToString(service.md5h.Sum(nil))
	data["songId"] = "0"
	data["version"] = "1"
	dataBytes, err := json.MarshalIndent(data, "", "")
	fmt.Print("dataBytes->" + string(dataBytes))
	options := &_Options{
		cookies: service.Cookies,
		crypto:  weapi,
		ua:      pc,
	}
	res, _ := responseDefault("POST", `https://interface.music.163.com/api/cloud/upload/check`, data, options)
	fmt.Print("res->" + res)
	if err != nil {
		panic(err)
	}
	data = make(map[string]interface{}, 0)
	data["bucket"] = ""
	data["ext"] = "mp3"
	data["filename"] = strings.Split(stat.Name(), ".mp3")[0]
	data["local"] = "false"
	data["nos_product"] = "3"
	data["type"] = "audio"
	data["md5"] = hex.EncodeToString(service.md5h.Sum(nil))
	tokenRes, _ := responseDefault("POST", `https://music.163.com/weapi/nos/token/alloc`, data, options)
	fmt.Print("33333" + tokenRes)
	var _tokenRes TokenRes
	//marshalIndent, err := json.MarshalIndent(tokenRes, "", "")
	str := []byte(tokenRes)
	err = json.Unmarshal(str, &_tokenRes)
	if err != nil {
		panic(err)
	}
	needUpload := gjson.Get(res, "needUpload")
	if needUpload.Value() == true {
		fmt.Println("上传")
		service.UploadPlugin(file)
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	fmt.Print("tokenRes->" + tokenRes)

	//// 上传封面
	//coverImgId, err := UploadImg(service.Cookies, config.Config.HomePath+"/tmp/img.png")
	data = make(map[string]interface{}, 0)
	songId := gjson.Get(res, "songId").Str
	data["md5"] = hex.EncodeToString(service.md5h.Sum(nil))
	data["songid"] = songId
	data["filename"] = stat.Name()
	data["song"] = service.SongMetadata.Title
	data["album"] = service.SongMetadata.Album
	//data["coverId"] = coverImgId
	data["artist"] = service.SongMetadata.Artist
	data["bitrate"] = "999000"
	data["resourceId"] = _tokenRes.Result.ResourceID
	res2Data, err := json.MarshalIndent(data, "", "")
	fmt.Print("dataRes->" + string(res2Data))
	res2, _ := responseDefault("POST", `https://music.163.com/api/upload/cloud/info/v2`, data, options)
	fmt.Print("res2->" + res2)
	data = make(map[string]interface{}, 0)
	songId = gjson.Get(res2, "songId").Str
	songId = strings.Replace(songId, ".", "", 1)
	fmt.Println("songId->", songId)
	data["songid"] = songId
	res3, _ := responseDefault("POST", `https://interface.music.163.com/api/cloud/pub/v2`, data, options)
	fmt.Print("res3->" + res3)

	return nil
}
func (service *UploadCloud) UploadPlugin(file *os.File) string {
	options := &_Options{
		cookies: service.Cookies,
		crypto:  weapi,
		ua:      pc,
	}
	data := make(map[string]interface{}, 1)
	data["bucket"] = "jd-musicrep-privatecloud-audio-public"
	data["ext"] = "mp3"
	data["filename"] = strings.Split(stat.Name(), ".mp3")[0]
	data["local"] = "false"
	data["nos_product"] = "3"
	data["type"] = "audio"
	data["md5"] = hex.EncodeToString(service.md5h.Sum(nil))
	dataReq, err := json.MarshalIndent(data, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Print("UploadPluginDataReq->" + string(dataReq))
	tokenRes, _ := responseDefault("POST", `https://music.163.com/weapi/nos/token/alloc`, data, options)
	var _tokenRes TokenRes
	str := []byte(tokenRes)
	err = json.Unmarshal(str, &_tokenRes)
	if err != nil {
		panic(err)
	}
	objectKey := strings.Replace(_tokenRes.Result.ObjectKey, "/", "%2F", -1)
	//objectKey := "obj%2FwoDDmMOBw6PClWzCnMK-%2F11823050297%2Fd2ad%2Fc904%2F3180%2F505185f21a50ea65ddc3a0f333bc21a5.mp3"
	headers := make(map[string]string)
	headers["x-nos-token"] = _tokenRes.Result.Token
	fmt.Println("_tokenRes.Result.Token->" + _tokenRes.Result.Token)
	headers["Content-MD5"] = hex.EncodeToString(service.md5h.Sum(nil))
	headers["Content-Type"] = "audio/mpeg"
	headers["Content-Length"] = strconv.FormatInt(stat.Size(), 10)
	//headers["ContentLength"] = strconv.FormatInt(stat.Size(),10)
	url := "http://45.127.129.8/jd-musicrep-privatecloud-audio-public/" + objectKey + "?offset=0&complete=true&version=1.0"
	uploadFile, _ := _UploadFile(url, service.FilePath, headers)
	fmt.Println("uploadFile->" + uploadFile)
	fmt.Println("uploadFile URL->" + url)
	return uploadFile
}
