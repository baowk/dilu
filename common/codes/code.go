package codes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

const (
	SUCCESS            = 200
	FAILURE            = 500
	AuthorizationError = 403
	NotFound           = 404
	NotLogin           = 401
	InvalidToken       = 555
	InvalidParameter   = 10000
	UserDoesNotExist   = 10001
	ServerError        = 10101
	TooManyRequests    = 10102
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
