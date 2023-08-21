package service

import (
	"encoding/json"
	"fmt"

	"github.com/baowk/dilu-core/common/utils/http_util"
)

var (
	appId  = "wx45726bf2e3d84409"
	secret = "90ce4450c3421a1d1e181dd45a256789"
)

// {
// 	"access_token":"ACCESS_TOKEN",
// 	"expires_in":7200,
// 	"refresh_token":"REFRESH_TOKEN",
// 	"openid":"OPENID",
// 	"scope":"SCOPE",
// 	"is_snapshotuser": 1,
// 	"unionid": "UNIONID"
//   }

type WechatLogin struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	OpenId         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
	Unionid        string `json:"unionid"`
}

func AuthWechat(code string) (string, error) {

	client := &(http_util.HTTPClient{
		BaseURL: "https://api.weixin.qq.com",
		Headers: map[string]string{
			"Content-Type": "application/json",
			//"Authorization": "Bearer xxxxxxxxxxxx",
		},
	})

	uri := "/sns/oauth2/access_token" //?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code"

	m := map[string]string{
		"appid":      appId,
		"secret":     secret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	bs, err := json.Marshal(m)
	if err != nil {

	}
	response, err := client.Post(uri, bs)
	if err != nil {
		fmt.Println("POST error:", err)
		return "", err
	}
	fmt.Println("POST response:", string(response))
	return string(response), nil
}
