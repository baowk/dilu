package codes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

const (
	SUCCESS            = 200   //成功
	FAILURE            = 500   //失败
	AuthorizationError = 403   //没有权限
	NotFound           = 404   //资源不存在
	NotLogin           = 401   //没有登录
	InvalidToken       = 555   //无效Token
	InvalidParameter   = 10000 //参数错误
	UserDoesNotExist   = 10001 //服务器错误
	ServerError        = 10101 //服务器错误
	RequestsFrequently = 10102 //请求太频繁
	CaptchaErr         = 10103 //获取验证码失败
)

const (
	LANG_ZH_CN = "zh-CN"
	LANG_ZH    = "zh"
	LANG_EN    = "en"
)

type Code struct {
	EnableI18N bool
	Lang       string
}

func (e *Code) Enable() bool {
	return e.EnableI18N
}

func (e *Code) DefLang() string {
	return e.Lang
}

func (e *Code) GetMsg(code int, c *gin.Context) string {
	var lang string
	if e.EnableI18N {
		acceptLanguate := c.GetHeader("Accept-Language")
		tags, _, _ := language.ParseAcceptLanguage(acceptLanguate)
		if len(tags) > 0 {
			lang = tags[0].String()
		}
	} else {
		lang = e.Lang
	}

	var ok bool
	var str string
	switch lang {
	case LANG_ZH_CN, LANG_ZH:
		str, ok = zhCNText[code]
	case LANG_EN:
		str, ok = enUSText[code]
	default:
		str, ok = zhCNText[code]
	}
	if !ok {
		return "unknown error"
	}
	return str
}
