package netease

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"wy_music_cloud/utils"
)

type Options struct {
	Crypto  string
	Ua      string
	Cookies []*http.Cookie
	Token   string
	Url     string
}

var UserAgentList = []string{
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
}

func chooseUserAgent(ua string) string {

	rand.Seed(time.Now().UnixNano())
	index := 0
	if ua == "" {
		index = rand.Intn(len(UserAgentList))
	} else if ua == "mobile" {
		index = rand.Intn(8)
	} else {
		index = rand.Intn(7) + 7
	}
	return UserAgentList[index]
}

func CreateRequest(method string, url string, data map[string]string, options *Options, headers map[string]string) (map[string]interface{}, []*http.Cookie) {
	req := requests.Requests()
	req.Header.Set("User-Agent", chooseUserAgent(options.Ua))
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	csrfToken := ""
	music_U := ""
	answer := map[string]interface{}{}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if strings.Contains(url, "music.163.com") {
		req.Header.Set("Referer", "https://music.163.com")
	}
	if options.Cookies != nil {
		for _, cookie := range options.Cookies {
			req.SetCookie(cookie)
			if cookie.Name == "__csrf" {
				csrfToken = cookie.Value
			}
			if cookie.Name == "MUSIC_U" {
				music_U = cookie.Value
			}
		}
	}
	if options.Crypto == "weapi" {
		data["csrf_token"] = csrfToken
		data = utils.Weapi(data)
		reg, _ := regexp.Compile(`/\w*api/`)
		url = reg.ReplaceAllString(url, "/weapi/")
	} else if options.Crypto == "linuxapi" {
		linuxApiData := make(map[string]interface{}, 3)
		linuxApiData["method"] = method
		reg, _ := regexp.Compile(`/\w*api/`)
		linuxApiData["url"] = reg.ReplaceAllString(url, "/api/")
		linuxApiData["params"] = data
		data = utils.Linuxapi(linuxApiData)
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
		url = "https://music.163.com/api/linux/forward"
	} else if options.Crypto == "eapi" {
		eapiData := make(map[string]interface{})
		for key, value := range data {
			eapiData[key] = value
		}
		rand.Seed(time.Now().UnixNano())
		header := map[string]string{
			"osver":       "",
			"deviceId":    "",
			"mobilename":  "",
			"appver":      "6.1.1",
			"versioncode": "140",
			"buildver":    strconv.FormatInt(time.Now().Unix(), 10),
			"resolution":  "1920x1080",
			"os":          "android",
			"channel":     "",
			"requestId":   strconv.FormatInt(time.Now().Unix()*1000, 10) + strconv.Itoa(rand.Intn(1000)),
			"MUSIC_U":     music_U,
		}

		for key, value := range header {
			req.SetCookie(&http.Cookie{Name: key, Value: value, Path: "/"})
		}
		eapiData["header"] = header
		data = utils.Eapi(options.Url, eapiData)
		reg, _ := regexp.Compile(`/\w*api/`)
		url = reg.ReplaceAllString(url, "/eapi/")
	}
	var resp *requests.Response
	var err error
	if method == "POST" {
		var form requests.Datas = data
		resp, err = req.Post(url, form)
	} else {
		resp, err = req.Get(url)
	}

	if err != nil {
		answer["code"] = 520
		answer["err"] = err.Error()
		return answer, nil
	}
	cookies := resp.Cookies()

	body := resp.Content()
	//fmt.Println(string(body))
	b := bytes.NewReader(body)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	// 数据被压缩 进行解码
	if err == nil {
		io.Copy(&out, r)
		body = out.Bytes()
	}

	err = json.Unmarshal(body, &answer)
	// 出错说明不是json
	if err != nil {
		//fmt.Println(string(body))
		// 可能是纯页面
		if strings.Index(string(body), "<!DOCTYPE html>") != -1 {
			answer["code"] = 200
			answer["html"] = string(body)
			return answer, cookies
		}
		answer["code"] = 500
		answer["err"] = err.Error()
		fmt.Println(string(body))
		return answer, nil
	}
	if _, ok := answer["code"]; !ok {
		answer["code"] = 200
	}
	indent, err := json.MarshalIndent(answer, "", "")
	fmt.Println("CreateRequest" + string(indent))
	return answer, cookies
}

func UploadFile(filePath string, url string, headers map[string]string) string {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	fileWrite, err := bodyWrite.CreateFormFile("songFile", stat.Name())
	fmt.Println("file.Name->" + stat.Name())
	_, err = io.Copy(fileWrite, file)
	if err != nil {
		fmt.Println("io Copy error", err)
		panic(err)
	}

	contentType := "audio/mpeg"
	bodyWrite.Close() //正确位置            ✓
	request, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		fmt.Println("http newrequest error", err)
		panic(err)
	}
	request.Header.Set("Content-Type", contentType)
	//request.Header.Set("User-Agent", chooseUserAgent(""))
	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}

	fmt.Println("ContentLength->" + strconv.FormatInt(stat.Size(), 10))

	resp, err := http.DefaultClient.Do(request)
	//bodyWrite.Close() //错误位置         ✘
	//在上传时，会提示真实长度和提交的长度不符合
	if err != nil {
		panic(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(respBytes)
}

func _UploadFile(url string, filePath string, headers map[string]string) (string, error) {
	var apiResponse struct {
		DownloadUrl string `json:"download_url"`
	}
	file, err := os.Open(filePath)

	//var body bytes.Buffer
	//
	//w := multipart.NewWriter(&body)
	//
	//wr, err := w.CreateFormFile("file",file.Name())
	//if err != nil {
	//	return "", err
	//}
	//
	//n, err := io.Copy(wr, file)
	//if err != nil {
	//	return "", err
	//}
	//
	//w.Close()
	readFile, err := ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(readFile)))
	if err != nil {
		return "", err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	info, err := file.Stat()
	req.Header.Add("Content-Length", fmt.Sprintf("%d", info.Size()))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	mjson, _ := json.Marshal(headers)
	mString := string(mjson)
	fmt.Printf("print mString:%s", mString)
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("can't close response body ", err)
		}
	}()

	if res.StatusCode == http.StatusConflict {
		return "domain.CacheURL + path", nil
	}
	respBytes, err := ioutil.ReadAll(res.Body)
	fmt.Println("上传结果->" + string(respBytes))
	defer file.Close()
	return apiResponse.DownloadUrl, nil
}
