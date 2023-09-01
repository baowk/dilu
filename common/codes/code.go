package codes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

const (
	SUCCESS             = 200   //成功
	FAILURE             = 500   //失败
	AuthorizationError  = 403   //没有权限
	NotFound            = 404   //资源不存在
	NotLogin            = 401   //没有登录
	InvalidToken        = 555   //无效Token
	ErrToken            = 1001  //Token失效
	ErrNoToken          = 1023  //登录用户才能访问
	ErrNoAccessRights   = 1024  //没有访问权限
	ErrUseLimitUser     = 1026  //您不在访问授权时间内
	ErrUserNotExist     = 1027  //当前账号不存在，请先注册
	ErrUserLock         = 1028  //当前账号被冻结
	ErrPwdNotExist      = 1029  //尚未设置密码，请选择验证码登录
	ErrMealLimit        = 1030  //成员数已达到套餐的限制，请升级套餐重试
	ErrAccessApplyExits = 1031  //您已经提交了登录申请，请耐心等待上级授权，通过授权即可登录
	ErrMpNotScan        = 1055  //未收到扫描信息
	ErrMpExpire         = 1056  //已过期
	InvalidParameter    = 10000 //参数错误
	ErrRePassword       = 10001 //重复密码不一致
	ErrPasswordFMT      = 10002 //密码长度必须6位
	ErrMobileOrEmail    = 10003 //必须手机号或者邮箱注册
	ErrParams           = 10004 //参数错误
	ErrVerifyCode       = 10005 //验证码错误
	ErrUnLogin          = 10006 //未登录
	ErrNotSelectCompany = 10007 //未选中企业
	ErrBind             = 10008 //绑定失败
	ErrUserExist        = 10011 //账号已经注册，请直接登录
	ErrRoleOnePerm      = 10012 //至少选择一项权限
	RequestsFrequently  = 10102 //请求太频繁
	CaptchaErr          = 10103 //获取验证码失败
	CaptchaVerifyErr    = 10104 //验证码
	PhoneExistErr       = 10201 //手机号已存在
	EmailExistErr       = 10202 //邮箱已存在
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
