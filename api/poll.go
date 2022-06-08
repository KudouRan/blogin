package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/imroc/req/v3"
)

type BaseResponse2 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
}

type QrLoginData struct {
	Mid          int    `json:"mid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type QrLogin struct {
	*BaseResponse2
	Data QrLoginData `json:"data"`
}

func (l QrLogin) ToJson() []byte {
	b, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return b
}

func getClient2() *req.Client {
	return req.C().SetCommonHeaders(map[string]string{
		"user-agent":   "Mozilla/5.0 BiliDroid/6.74.0 (bbcallen@gmail.com) os/android model/MI 10 Pro mobi_app/android build/6740400 channel/xiaomi innerVer/xiaomi osVer/10 network/2",
		"content-type": "application/x-www-form-urlencoded; charset=utf-8",
	})
}

/**
 * app sign 签名
 * @param params *map[string]string
 */
func appSign2(params map[string]string) string {
	appkey := "4409e2ce8ffd12b8"
	appsec := "59b43e04ad6965f34319062b478f83dd"
	// 给 params 加入 appkey
	(params)["appkey"] = appkey
	// 对 params 进行排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 将 params 编码成 query string
	var query string
	for _, k := range keys {
		query += fmt.Sprintf("%s=%s&", k, url.QueryEscape((params)[k]))
	}
	// 对 query 进行 md5
	hash := md5.New()
	hash.Write([]byte(query[:len(query)-1] + appsec))
	(params)["sign"] = hex.EncodeToString(hash.Sum(nil))
	return query + "sign=" + (params)["sign"]
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	client := getClient2()
	api := "https://passport.bilibili.com/x/passport-tv-login/qrcode/poll"
	data := appSign2(map[string]string{
		"ts":       strconv.FormatInt(time.Now().Unix(), 10),
		"local_id": "0",
	})
	var qrBody QrLogin
	resp, _ := client.R().SetBody(data).SetResult(&qrBody).Post(api)
	if resp.IsError() {
		w.Write([]byte("Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	// 将结构体转换成 json 字符串返回
	w.Write(qrBody.ToJson())
}
