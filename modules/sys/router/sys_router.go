package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckRoleRouter)
}

// 无需认证的路由示例
func sysNoCheckRoleRouter(v1 *gin.RouterGroup) {
	r := v1.Group("")
	{
		r.GET("captcha", apis.ApiCaptcha.GenerateCaptchaHandler)
	}
}
