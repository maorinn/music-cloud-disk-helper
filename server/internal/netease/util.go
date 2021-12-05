package netease

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func queryParam(data interface{}) (res []string) {
	res = make([]string, 0)
	m := data.(map[string]interface{})
	if m == nil {
		return res
	}
	for k, v := range m {
		if vm, ok := v.(map[string]interface{}); ok {
			param := queryParam(vm)
			res = append(res, param...)
			break
		}
		if vs, ok := v.(string); ok {
			res = append(res, k+"="+url.QueryEscape(vs))
		}
	}
	return
}

func queryParamString(data interface{}) string {
	param := queryParam(data)
	join := strings.Join(param, "&")
	return join
}

func getCookie(cookies []*http.Cookie, name string, defaultValue ...string) string {
	for _, v := range cookies {
		if v.Name == name {
			return v.Value
		}
	}
	if defaultValue == nil || len(defaultValue) == 0 {
		return ""
	}
	return defaultValue[0]
}

// ReadFile 读取文件
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []byte{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	fileContent, err := ioutil.ReadAll(file)
	return fileContent, err
}
