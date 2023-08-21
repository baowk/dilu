package service

// import (
// 	"dilu/modules/sys/service/dto"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"

// 	"github.com/faabiosr/cachego/file"

// 	"github.com/fastwego/dingding"
// )

// var DingClient *dingding.Client
// var DingConfig map[string]string

// var (
// 	agent_id   = "2464941691"
// 	app_key    = "dingc7p3yu5c74klxobd"
// 	app_secret = "Xu8qjSAu9XCYqaZmpEonvacFuHzmPY0KmFCFurhMdpjGbfPx9KmViKpaSANAGi27"
// 	crop_id    = ""
// )

// func init() {
// 	// 加载配置文件
// 	// viper.SetConfigFile(".env")
// 	// _ = viper.ReadInConfig()

// 	DingConfig = map[string]string{
// 		"CorpId":    crop_id,
// 		"AgentId":   agent_id,
// 		"AppKey":    app_key,
// 		"AppSecret": app_secret,
// 		// "Token":          viper.GetString("TOKEN"),
// 		// "EncodingAESKey": viper.GetString("EncodingAESKey"),
// 	}

// 	// 钉钉 AccessToken 管理器
// 	atm := &dingding.DefaultAccessTokenManager{
// 		Id:   DingConfig["AppKey"],
// 		Name: "access_token",
// 		GetRefreshRequestFunc: func() *http.Request {
// 			params := url.Values{}
// 			params.Add("appkey", DingConfig["AppKey"])
// 			params.Add("appsecret", DingConfig["AppSecret"])
// 			req, _ := http.NewRequest(http.MethodGet, dingding.ServerUrl+"/gettoken?"+params.Encode(), nil)

// 			return req
// 		},
// 		Cache: file.New(os.TempDir()),
// 	}

// 	// 钉钉 客户端
// 	DingClient = dingding.NewClient(atm)

// 	atm.Cache = file.New(os.TempDir())
// }

// type DingUser struct {
// 	Userid   string `json:"userid"`
// 	SysLevel int    `json:"sys_level"`
// 	Errmsg   string `json:"errmsg"`
// 	IsSys    bool   `json:"is_sys"`
// 	Errcode  int    `json:"errcode"`
// }

// func GetDingConfig(domain string, cfg *dto.DingCfgResp) {
// 	state := utils.GenUUid()
// 	cfg.Appid = DingConfig["AppKey"]
// 	cfg.State = state
// 	cfg.RedirectUri = "https://www.yunlogin.com/login"
// }

// func AuthDing(code string) ([]byte, error) {
// 	fmt.Printf("2222")
// 	var userInfo []byte
// 	fmt.Println("code = ", code)
// 	if len(code) == 0 {
// 		return userInfo, errors.New("auth code is null")
// 	}

// 	// 获取用户身份
// 	params := url.Values{}
// 	params.Add("code", code)

// 	req, _ := http.NewRequest(http.MethodGet, "/user/getuserinfo?"+params.Encode(), nil)
// 	userInfo, err := DingClient.Do(req)
// 	log.Println(string(userInfo), err)
// 	if err != nil {
// 		return userInfo, errors.New("auth code is null")
// 	}

// 	return userInfo, nil
// }
