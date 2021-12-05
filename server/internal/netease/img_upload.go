package netease

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func UploadImg(cookies []*http.Cookie, filePath string) (string, error) {
	cookies = append(cookies, &http.Cookie{Name: "os", Value: "pc"}, &http.Cookie{Name: "appver", Value: "2.7.1.198277"})
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	options := &Options{
		Crypto:  "weapi",
		Ua:      "pc",
		Cookies: cookies,
	}
	defer file.Close()
	info, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	// 获取key和token
	data := map[string]string{
		"bucket":      "yyimgs",
		"ext":         "jpg",
		"filename":    info.Name(),
		"local":       "false",
		"nos_product": "0",
		"return_body": `{"code":200,"size":"$(ObjectSize)"}`,
		"type":        "other",
	}
	tokenRes, _ := CreateRequest("POST", `https://music.163.com/weapi/nos/token/alloc`, data, options, nil)
	var _tokenRes TokenRes
	marshalIndent, err := json.MarshalIndent(tokenRes, "", "")
	err = json.Unmarshal(marshalIndent, &_tokenRes)
	if err != nil {
		panic(err)
	}
	// 上传图片
	uploadFile := UploadFile(filePath, "https://nosup-hz1.127.net/yyimgs/"+_tokenRes.Result.ObjectKey+"?offset=0&complete=true&version=1.0", map[string]string{
		"x-nos-token":  _tokenRes.Result.Token,
		"Content-Type": "image/jpeg",
	})
	fmt.Printf("上传图片1111111111RES%s,图片id%s", uploadFile, _tokenRes.Result.DocID)
	return _tokenRes.Result.DocID, nil
}
