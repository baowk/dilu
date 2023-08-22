package codes

var zhCNText = map[int]string{
	SUCCESS:            "OK",
	FAILURE:            "FAIL",
	NotFound:           "资源不存在",
	ServerError:        "服务器内部错误",
	RequestsFrequently: "请求过多",
	InvalidParameter:   "参数错误",
	UserDoesNotExist:   "用户不存在",
	AuthorizationError: "暂无访问权限",
	NotLogin:           "请先登录",
	InvalidToken:       "无效的token",
	CaptchaErr:         "获取验证码失败，请稍后重试.",
}
