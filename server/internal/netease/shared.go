package netease

import (
	"io"
	"net/http"
	"net/url"
)

const (
	mobile   userAgentType = "mobile"
	pc       userAgentType = "pc"
	weapi    cryptoType    = "weapi"
	eapi     cryptoType    = "eapi"
	linuxapi cryptoType    = "linuxapi"
)

type userAgentType string

type cryptoType string

type Query struct {
	Cookies []*http.Cookie
	Param   url.Values
	Body    io.ReadCloser
	Proxy   interface{}
}

func (query *Query) AddParam(name, value string) *Query {
	if query.Param == nil {
		query.Param = url.Values{}
	}
	query.Param.Add(name, value)
	return query
}

func (query *Query) GetParam(name string) string {
	if query.Param == nil {
		return ""
	}
	return query.Param.Get(name)
}

func (query *Query) GetParamOrDefault(name string, defaultValue ...interface{}) interface{} {
	var dv interface{}
	if defaultValue != nil && len(defaultValue) > 0 {
		dv = defaultValue[0]
	}
	if query.Param == nil {
		return dv
	}
	if p := query.Param.Get(name); p != "" {
		return p
	}
	return dv
}

func (query *Query) AddCookie(name, value string) *Query {
	if query.Cookies == nil {
		query.Cookies = make([]*http.Cookie, 0)
	}
	query.Cookies = append(query.Cookies, &http.Cookie{
		Name:  name,
		Value: value,
	})
	return query
}

func (query *Query) GetCookie(name string) string {
	if query.Cookies == nil {
		return ""
	}
	return getCookie(query.Cookies, name)
}

type _Options struct {
	crypto  cryptoType
	cookies []*http.Cookie
	proxy   interface{}
	ua      userAgentType
	token   string
	url     string
}
