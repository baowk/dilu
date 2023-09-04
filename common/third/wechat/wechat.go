package wechat

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/baowk/dilu-core/common/utils/https"
)

const (
	redirectOauthURL       = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	webAppRedirectOauthURL = "https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	webAccessTokenURL      = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL  = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	userInfoURL            = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	checkAccessTokenURL    = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	accessTokenURL         = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	getTicketURL           = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	getQrCode              = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	getMpUserInfo          = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
	//获取二维码地址 :https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=gQF88DwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAySFJ1NFlQUGtjd0MxWjNKQzFBY08AAgQX7OZkAwQsAQAA
)

// 获取redirectOauthURL链接。其在微信中跳转后可以获取code
// 因为微信转码会造成部分链接参数丢失的情况，使用urlEncode对链接进行处理
func RedirectOauthUrl(appID, redirectUrl string) string {
	if appID == "" || redirectUrl == "" {
		return ""
	}

	// url encode
	v := url.Values{}
	v.Add("redirectUrl", redirectUrl) // 添加map
	encodeUrl := v.Encode()
	encodeUrl = strings.TrimLeft(encodeUrl, "redirectUrl=") //去掉url中多余的字符串
	urlStr := fmt.Sprintf(redirectOauthURL, appID, encodeUrl, "snsapi_userinfo", "")
	return urlStr
}

// 网页授权access_token
type ResWebAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Openid       string `json:"openid"`
	Unionid      string `json:"unionid"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

// 获取网页授权access_token
func GetWebAccessToken(appID, secret, code string) (res ResWebAccessToken, err error) {
	urlStr := fmt.Sprintf(webAccessTokenURL, appID, secret, code)
	body, err := https.New().Get(urlStr)
	fmt.Println("GetWebAccessToken:" + string(body))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("GetWebAccessToken error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}
	return
}

type WxUserInfo struct {
	ID         int        `json:"id"`
	Openid     string     `json:"openid"`
	Nickname   string     `json:"nickname"`
	Headimgurl string     `json:"headimgurl"`
	Sex        int        `json:"sex"`
	Province   string     `json:"province"`
	City       string     `json:"city"`
	Country    string     `json:"country"`
	Name       string     `json:"name"`
	Mobile     string     `json:"mobile"`
	Address    string     `json:"address"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Errcode    int        `json:"errcode"`
	Errmsg     string     `json:"errmsg"`
}

func GetUserInfo(accessToken, openID string) (res WxUserInfo, err error) {
	urlStr := fmt.Sprintf(userInfoURL, accessToken, openID)
	body, err := https.New().Get(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}
	return
}

// 普通access_token
type ResAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

// 获取普通access_token
func GetAccessToken(appID, secret string) (res ResAccessToken, err error) {
	urlStr := fmt.Sprintf(accessTokenURL, appID, secret)
	body, err := https.New().Get(urlStr)
	fmt.Println("[GetAccessToken]:" + string(body))
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}
	return
}

type resTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
}

func GetTicket(accessToken string) (res resTicket, err error) {
	urlStr := fmt.Sprintf(getTicketURL, accessToken)
	body, err := https.New().Get(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	if res.Errcode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", res.Errcode, res.Errmsg)
		return
	}
	return
}

// {"ticket":"gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm
// 3sUw==","expire_seconds":60,"url":"http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"}
type QrCodeResp struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
	Scene         string `json:"scene"`
}

// {"expire_seconds": 604800, "action_name": "QR_SCENE", "action_info": {"scene": {"scene_id": 123}}} 或者也可以使用以下POST数据创建字符串形式的二维码参数：{"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
// {"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_id": 123}}} 或者也可以使用以下POST数据创建字符串形式的二维码参数： {"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
type QrCodeReq struct {
	ExpireSeconds int            `json:"expire_seconds,omitempty"` //该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为60秒。
	ActionName    string         `json:"action_name"`              //二维码类型，QR_SCENE为临时的整型参数值，QR_STR_SCENE为临时的字符串参数值，QR_LIMIT_SCENE为永久的整型参数值，QR_LIMIT_STR_SCENE为永久的字符串参数值
	Actioninfo    map[string]any `json:"action_info"`
}

// 获取二维码 expire 大于0 二维码有效期 秒
func GetQrCode(accessToken string, expire int, sceneId int, sceneStr string) (res QrCodeResp, err error) {
	urlStr := fmt.Sprintf(getQrCode, accessToken)
	m := make(map[string]any, 0)
	p := QrCodeReq{}
	if expire > 0 {
		p.ExpireSeconds = expire
		if sceneId == 0 {
			p.ActionName = "QR_STR_SCENE"
			m["scene_str"] = sceneStr
		} else {
			p.ActionName = "QR_SCENE"
			m["scene_id"] = sceneId
		}
	} else {
		if sceneId == 0 {
			p.ActionName = "QR_LIMIT_STR_SCENE"
			m["scene_str"] = sceneStr
		} else {
			p.ActionName = "QR_LIMIT_SCENE"
			m["scene_id"] = sceneId
		}
	}
	data := make(map[string]any, 0)
	data["scene"] = m
	p.Actioninfo = data
	//p.Actioninfo = m

	bs, err := json.Marshal(p)
	if err != nil {
		return
	}
	body, err := https.New().Post(urlStr, bs)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	return
}

type WxSignCheck struct {
}

func MPCheckSign(timestamp, nonce, signature, wxToken string) bool {
	params := sort.StringSlice{timestamp, nonce, wxToken}
	sort.Sort(params)
	s := strings.Join(params, "")
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
	//fmt.Printf("signature:%s \n", signature)
	res := hex.EncodeToString(h.Sum(nil))
	//fmt.Printf("mysign   :%s \n", res)
	return res == signature
}

type MPMsg struct {
	ToUserName   string `xml:"ToUserName"`   //开发者微信号
	FromUserName string `xml:"FromUserName"` //发送方账号（一个OpenID）
	CreateTime   int    `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType      string `xml:"MsgType"`      //消息类型，event
	Event        string `xml:"Event"`        //事件类型，scancode_push
	EventKey     string `xml:"EventKey"`     //事件KEY值，由开发者在创建菜单时设定
	ScanCodeInfo string `xml:"ScanCodeInfo"` //扫描信息
	ScanType     string `xml:"ScanType"`     //扫描类型，一般是qrcode
	ScanResult   string `xml:"ScanResult"`   //扫描结果，即二维码对应的字符串信息
}

type MpUserinfo struct {
	Subscribe      int    `json:"subscribe"`       // 1, 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenId         string `json:"openid"`          // "o6_bmjrPTlm6_2sgVt7hMZOPfL2M", //用户的标识，对当前公众号唯一
	Language       string `json:"language"`        // "zh_CN", //用户的语言，简体中文为zh_CN
	SubscribeTime  int    `json:"subscribe_time"`  //1382694957,//用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid        string `json:"unionid"`         // " o6_bmasdasdsad6_2sgVt7hMZOPfL",//只有在用户将公众号绑定到微信开放平台账号后，才会出现该字段。
	Remark         string `json:"remark"`          // "",//公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid        int    `json:"groupid"`         //0,//用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int  `json:"tagid_list"`      //[128,2],//用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene"` // "ADD_SCENE_QR_CODE",//返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_REPRINT 他人转载 ,ADD_SCENE_LIVESTREAM 视频号直播，ADD_SCENE_CHANNELS 视频号 , ADD_SCENE_OTHERS 其他
	QrScene        int    `json:"qr_scene"`        //: 98765,二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str"`    //: ""二维码扫码场景描述（开发者自定义）
}

func GetMpUserInfo(accessToken, openId string, ui *MpUserinfo) error {
	urlStr := fmt.Sprintf(getMpUserInfo, accessToken, openId)
	body, err := https.New().Get(urlStr)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return json.Unmarshal(body, ui)
}
