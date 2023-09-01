package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckSsoRouter)
}

// 无需认证的路由示例
func sysNoCheckSsoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sso")
	{
		r.POST("login", apis.Sso.Login)
		r.POST("register", apis.Sso.Register)
		r.POST("sendCode", apis.Sso.SendCode)
		r.POST("forgetPwd", apis.Sso.ForgetPwd)
	}
}
