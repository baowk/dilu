package codes

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

func GetMsg(code int, lang string) (str string) {
	var ok bool
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
	return
}
