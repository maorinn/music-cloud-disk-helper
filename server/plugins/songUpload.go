package plugins

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dhowden/tag"
	"hash"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"wy_music_cloud/utils"
	//"github.com/dhowden/tag"
)
type TokenRes struct {
	Code int `json:"code"`
	Message interface{} `json:"message"`
	Result Result `json:"result"`
}
type Result struct {
	Bucket string `json:"bucket"`
	DocID string `json:"docId"`
	ObjectKey string `json:"objectKey"`
	OuterURL string `json:"outerUrl"`
	ResourceID int64 `json:"resourceId"`
	Token string `json:"token"`
}

type UploadCloud struct {
	FilePath string
	Cookies []*http.Cookie
	md5h hash.Hash
}

var (
	songName = ""
	artist = "未知艺术家"
	album = "未知专辑"
	stat os.FileInfo
)
func (service *UploadCloud) UploadCloud() map[string]interface{} {
	service.Cookies = append(service.Cookies,&http.Cookie{Name: "os", Value: "pc"},&http.Cookie{Name: "appver",Value: "2.7.1.198277"})
	file, err := os.Open(service.FilePath)


	defer file.Close()
	service.md5h =  md5.New()
	io.Copy(service.md5h, file)
	if err != nil {
		panic(err)
	}
	stat, err = os.Stat(service.FilePath)

	if err != nil {
		panic(err)
	}
	songName =  strings.Split(stat.Name(),".mp3")[0]
	fmt.Print(stat.Size())
	data := make(map[string]string)
	data["bitrate"] = "999000"
	data["ext"] = ""
	data["length"] = strconv.FormatInt(stat.Size(),10)
	data["md5"] = 	hex.EncodeToString(service.md5h.Sum(nil))
	data["songId"] = "0"
	data["version"] = "1"
	dataBytes, err := json.MarshalIndent(data, "", "")
	fmt.Print("dataBytes->"+string(dataBytes))
	options := &utils.Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: service.Cookies,
	}
	res, _ := utils.CreateRequest("POST", `https://interface.music.163.com/api/cloud/upload/check`, data, options,nil);
	indent, err := json.MarshalIndent(res, "", "")
	fmt.Print("111"+string(indent))
	m, err := tag.ReadFrom(file)
	//if err != nil {
	//	panic(err)
	//}
	if m != nil {
		if m.Title() != "" {
			songName = m.Title()
		}
		if m.Artist()!="" {
			artist = m.Artist()
		}
		if m.Album()!="" {
			album = m.Album()
		}
	}

	data = make(map[string]string)
	data["bucket"] = ""
	data["ext"] = "mp3"
	data["filename"] = strings.Split(stat.Name(),".mp3")[0]
	data["local"] = "false"
	data["nos_product"] = "3"
	data["type"] = "audio"
	data["md5"] = hex.EncodeToString(service.md5h.Sum(nil))
	tokenRes, _ := utils.CreateRequest("POST", `https://music.163.com/weapi/nos/token/alloc`, data, options,nil);
	var _tokenRes TokenRes
	marshalIndent, err := json.MarshalIndent(tokenRes, "", "")
	err = json.Unmarshal(marshalIndent, &_tokenRes)
	if err != nil {
		panic(err)
	}
	if res["needUpload"]==true {
		fmt.Println("上传")
		service.UploadPlugin(file)
	}
	fmt.Print("tokenRes->"+string(marshalIndent))

	data = make(map[string]string)
	data["songid"] = res["songId"].(string)
	data["song"] = songName
	data["filename"] = strings.Split(stat.Name(),".mp3")[0]
	data["album"] = album
	data["artist"] = artist
	data["bitrate"] = "999000"
	data["md5"] = hex.EncodeToString(service.md5h.Sum(nil))
	data["resourceId"] =  strconv.FormatInt(_tokenRes.Result.ResourceID,10)
	dataRes, err := json.MarshalIndent(data, "", "")

	fmt.Print("dataRes->"+string(dataRes))
	res2, _ := utils.CreateRequest("POST", `https://music.163.com/api/upload/cloud/info/v2`, data, options,nil)
	marshalIndent, err = json.MarshalIndent(res2, "", "")
	fmt.Print("res2->"+string(marshalIndent))
	data = make(map[string]string)
	data["songid"] = res2["songId"].(string)
	res3, _ := utils.CreateRequest("POST", `https://interface.music.163.com/api/cloud/pub/v2`, data, options,nil)
	marshalIndent, err = json.MarshalIndent(res3, "", "")
	fmt.Print("res2->"+string(marshalIndent))

	return nil
}
func (service *UploadCloud) UploadPlugin(file *os.File) string {
	options := &utils.Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: service.Cookies,
	}
	data := make(map[string]string)
	data["bucket"] = "jd-musicrep-privatecloud-audio-public"
	data["ext"] = "mp3"
	data["filename"] = strings.Split(stat.Name(),".mp3")[0]
	data["local"] = "false"
	data["nos_product"] = "3"
	data["type"] = "audio"
	data["md5"] = hex.EncodeToString(service.md5h.Sum(nil))
	dataReq, err := json.MarshalIndent(data, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Print("UploadPluginDataReq->"+string(dataReq))
	tokenRes, _ := utils.CreateRequest("POST", `https://music.163.com/weapi/nos/token/alloc`, data, options,nil);
	var _tokenRes TokenRes
	marshalIndent, err := json.MarshalIndent(tokenRes, "", "")
	fmt.Print("UploadPluginRes-->"+string(marshalIndent))
	err = json.Unmarshal(marshalIndent, &_tokenRes)
	if err != nil {
		panic(err)
	}
	objectKey := strings.Replace(_tokenRes.Result.ObjectKey,"/","%2F",-1)
	//objectKey := "obj%2FwoDDmMOBw6PClWzCnMK-%2F11823050297%2Fd2ad%2Fc904%2F3180%2F505185f21a50ea65ddc3a0f333bc21a5.mp3"
	headers := make(map[string]string)
	headers["x-nos-token"] = _tokenRes.Result.Token
	headers["Content-MD5"] = hex.EncodeToString(service.md5h.Sum(nil))
	headers["Content-Type"] = "audio/mpeg"
	headers["Content-Length"] = strconv.FormatInt(stat.Size(),10)
	//headers["ContentLength"] = strconv.FormatInt(stat.Size(),10)
	uploadFile := utils.UploadFile(service.FilePath, "http://45.127.129.8/jd-musicrep-privatecloud-audio-public/"+objectKey+"?offset=0&complete=true&version=1.0", headers)
	fmt.Print("cccccc"+uploadFile)
	return uploadFile
}
