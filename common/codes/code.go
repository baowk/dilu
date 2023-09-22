package codes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

const (
	SUCCESS                = 200   //成功
	FAILURE                = 500   //失败
	InvalidToken_401       = 401   //没有登录
	AuthorizationError_403 = 403   //没有权限
	NotFound_404           = 404   //资源不存在
	UserNotExist           = 1027  //当前账号不存在，请先注册
	UserLock               = 1028  //当前账号被冻结
	PwdNotExist            = 1029  //尚未设置密码，请选择验证码登录
	ThirdNotScan           = 1055  //未收到扫描信息
	ThirdExpire            = 1056  //已过期
	InvalidParameter       = 10000 //参数错误
	ErrRePassword          = 10001 //重复密码不一致
	ErrPasswordFMT         = 10002 //密码长度必须6位
	ErrMobileOrEmail       = 10003 //必须手机号或者邮箱注册
	ErrParams              = 10004 //参数错误
	ErrVerifyCode          = 10005 //验证码错误
	ErrBind                = 10008 //绑定失败
	ErrUserExist           = 10011 //账号已经注册，请直接登录
	ErrPwd                 = 10012 //密码错误
	ErrUsernameOrPwd       = 10013 //账号或密码错误
	RequestsFrequently     = 10102 //请求太频繁
	CaptchaErr             = 10103 //获取验证码失败
	CaptchaVerifyErr       = 10104 //验证码
	PhoneExistErr          = 10201 //手机号已存在
	EmailExistErr          = 10202 //邮箱已存在
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
