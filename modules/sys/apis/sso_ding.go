package apis

import (
	"crypto/hmac"
	"crypto/sha256"
	"dilu/common"
	"dilu/common/config"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Ding struct {
	base.BaseApi
}

// 获取钉钉登录配置信息
// GeDingCfg 获取钉钉登录配置信息
// @Summary 获取钉钉登录配置信息
// @Description 获取钉钉登录配置信息
// @Tags sys-sso ding
// @Success 200 {object} base.Resp{data=dto.DingCfgResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/getDingCfg [post]
func (e Ding) GetDingCfg(c *gin.Context) {

	//domain := "http://" + c.Request.Host
	//fmt.Println(domain)
	var cfg dto.DingCfgResp
	//service.GetDingConfig(domain, &cfg)
	e.Ok(c, cfg)
}

// 获取钉钉回调
// GeDingCfg 获取钉钉回调
// @Summary 获取钉钉回调
// @Description 获取钉钉回调
// @Tags sys-sso ding
// @Success 200 {object} base.Resp{data=dto.DingCfgResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/ding/callback [get]
func (e Ding) DingCallback(c *gin.Context) {

	code := c.Query("code")
	if code == "" {
		code = c.Query("loginTmpCode")
	}
	state := c.Query("state")

	uid, err := LoginByQRcode(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("uid:" + uid)
	common.SetMpOpenId("ding:"+state, uid)
}

// 钉钉
// LoginByDing 钉钉登录
// @Summary 钉钉
// @Description 钉钉登录
// @Tags sys-sso ding
// @Accept application/json
// @Product application/json
// @Param data body dto.LoginDingReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/loginDing [post]
func (e Ding) LoginByDing(c *gin.Context) {
	req := dto.LoginDingReq{}

	SnsAuthorize(req.State, req.Code)

	//ip := common.GetClientIP(c)
	userId, err := common.GetMpOpenId("ding:" + req.State)
	if err != nil {
		userId, err = LoginByQRcode(req.Code)
		if err != nil {
			core.Log.Error("ding login", zap.Error(err))
			e.Error(c, err)
			return
		}
	}

	if logOk, err := service.SerSysUser.LoginDing(&req, userId); err != nil {
		e.Error(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// 调用钉钉auth
func SnsAuthorize(state string, code string) {
	url := "https://oapi.dingtalk.com/connect/oauth2/sns_authorize?" +
		"appid=" + config.Ext.Ding.AppKey +
		"&response_type=code" +
		"&scope=snsapi_login" +
		"&state=" + state +
		"&redirect_uri=https://www.yunlogin.com/api/v1/sys/ding/callback" +
		"&loginTmpCode=" + code
	http.Get(url)
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//body, _ := io.ReadAll(resp.Body)
}

func LoginByQRcode(code string) (userid string, err error) {
	var resp *http.Response
	//fmt.Println("AppKey,AppSecret", AppKey, AppSecret)
	//服务端通过临时授权码获取授权用户的个人信息
	appKey := config.Ext.Ding.AppKey
	appSecret := config.Ext.Ding.AppSecret
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10) // 毫秒时间戳
	signature := EncodeSHA256(timestamp, appSecret)                   // 加密签名  加密算法见我另一个函数
	url2 := fmt.Sprintf(
		"https://oapi.dingtalk.com/sns/getuserinfo_bycode?accessKey=%s&timestamp=%s&signature=%s",
		appKey, timestamp, signature)
	//fmt.Println(3, url2, )
	p := struct {
		Tmp_auth_code string `json:"tmp_auth_code"`
	}{code} // post数据
	p1, _ := json.Marshal(p)
	p2 := string(p1)
	p3 := strings.NewReader(p2) //构建post数据
	resp, err = http.Post(url2, "application/json;charset=UTF-8", p3)
	if err != nil {
		return "", err
	}
	//fmt.Println(1, resp, err)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(2, string(body), err)
	var i map[string]interface{}
	_ = json.Unmarshal(body, &i) ///返回的数据给i
	errcode := i["errcode"].(float64)
	if errcode != 0 {
		return "", fmt.Errorf("登录错误: %f, %s", errcode, i["errmsg"].(string))
	}
	unionid := i["user_info"].(map[string]interface{})["unionid"].(string) // unionid 可以用来查询userinfo
	accesstoken, err := GetAccesstoken()                                   // 获取accesstoken
	if err != nil {
		return "", fmt.Errorf("登录错误accesstoken获取失败: %s", err.Error())
	}
	userid, err = GetUseridByUnionid(accesstoken, unionid)
	if err != nil {
		return "", fmt.Errorf("登录错误userid获取失败: %s", err)
	}
	return userid, nil
}

func GetUseridByUnionid(accesstoken, unionid string) (userid string, err error) {
	//根据unionid获取userid
	var resp *http.Response
	url := fmt.Sprintf("https://oapi.dingtalk.com/user/getUseridByUnionid?access_token=%s&unionid=%s",
		accesstoken, unionid)
	resp, err = http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(1, string(body), err)
	var i map[string]interface{}
	_ = json.Unmarshal(body, &i)
	errcode := i["errcode"].(float64)
	if errcode != 0 {
		return "", fmt.Errorf("userid获取错误: %f, %s", errcode, i["errmsg"].(string))
	}
	return i["userid"].(string), nil
}

func GetAccesstoken() (accesstoken string, err error) {
	var resp *http.Response
	//var AppKey, AppSecret string
	//获取access_token
	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", config.Ext.Ding.AppKey, config.Ext.Ding.AppSecret)
	resp, err = http.Get(url)
	if err != nil {
		return "", err
	}
	//fmt.Println(resp)
	//fmt.Println(err)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var i map[string]interface{}
	_ = json.Unmarshal(body, &i)
	//fmt.Println(1, string(body), i["errmsg"])
	if i["errcode"].(float64) == 0 {
		return i["access_token"].(string), nil
	}
	return "", errors.New("accesstoken获取错误：" + i["errmsg"].(string))

}

func EncodeSHA256(message, secret string) string {
	// 钉钉签名算法实现
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	sum := h.Sum(nil) // 二进制流
	message1 := base64.StdEncoding.EncodeToString(sum)

	uv := url.Values{}
	uv.Add("0", message1)
	message2 := uv.Encode()[2:]
	return message2

}

func DingTmpHtml(c *gin.Context) {
	t1, err := template.ParseFiles("app/sso/apis/ding.html")
	if err != nil {
		core.Log.Error("template err", zap.Error(err))
	}
	t1.Execute(c.Writer, "")
}
