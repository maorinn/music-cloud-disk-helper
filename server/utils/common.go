package utils

import (
	"math/rand"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func interfaceMapToString(_map map[string] interface{})map[string]string  {
	resultMap:= make(map[string]string)
	for key,value := range _map {
		resultMap[key] = value.(string)
	}
	return resultMap
}
