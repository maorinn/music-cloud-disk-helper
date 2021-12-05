package netease

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/io24m/hammer/crypto"
	aesPkg "github.com/io24m/hammer/crypto/aes"
	rsaPkg "github.com/io24m/hammer/crypto/rsa"
	"strings"
)

const (
	presetKey   = "0CoJUm6Qyw8W8jud"
	linuxapiKey = "rFgB&h#%2?^eDg:Q"
	ivParameter = "0102030405060708"
	publicKey   = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----"
)

func weapiEncrypt(data interface{}) (res map[string]interface{}) {
	res = make(map[string]interface{})
	jsonStr, _ := json.Marshal(data)
	secretKey := crypto.RKey(16)
	rKey := crypto.ReverseKey(secretKey)
	encrypt, _ := aesPkg.AesEncryptCBC(jsonStr, []byte(presetKey), []byte(ivParameter))
	b64 := base64Encode(encrypt)
	aes128Encrypt, _ := aesPkg.AesEncryptCBC([]byte(b64), secretKey, []byte(ivParameter))
	b64 = base64Encode(aes128Encrypt)
	res["params"] = string(b64)
	r, _ := rsaPkg.RsaEncryptNoPadding(rKey, []byte(publicKey))
	res["encSecKey"] = r
	return
}

func linuxapiEncrypt(data interface{}) (res map[string]interface{}) {
	res = make(map[string]interface{})
	jsondata, _ := json.Marshal(data)
	ecb, _ := aesPkg.AesEncryptECB(jsondata, []byte(linuxapiKey))
	res["eparams"] = strings.ToUpper(hex.EncodeToString(ecb))
	return
}

func eapiEncrypt(url string, data interface{}) (res map[string]interface{}) {
	text := ""
	s, ok := data.(string)
	if ok {
		text = s
	} else {
		jsonData, _ := json.Marshal(data)
		text = string(jsonData)
	}
	message := fmt.Sprintf("nobody%suse%smd5forencrypt", url, text)
	sum := md5.Sum([]byte(message))
	digest := hex.EncodeToString(sum[:])
	pa := fmt.Sprintf("%s-36cd479b6b5-%s-36cd479b6b5-%s", url, text, digest)
	res = make(map[string]interface{})
	ecb, _ := aesPkg.AesEncryptECB([]byte(pa), []byte(linuxapiKey))
	res["params"] = strings.ToUpper(hex.EncodeToString(ecb))
	return
}

func decrypt(data interface{}) interface{} {
	//aesDecryptECB()
	return nil
}

func base64Encode(data []byte) (buf []byte) {
	stdEncoding := base64.StdEncoding
	buf = make([]byte, stdEncoding.EncodedLen(len(data)))
	stdEncoding.Encode(buf, data)
	return
}
