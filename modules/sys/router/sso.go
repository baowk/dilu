package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckSsoRouter)
}

// 无需认证的路由示例
func sysNoCheckSsoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("")
	{
		r.POST("login", apis.ApiSso.Login)
		r.POST("register", apis.ApiSso.Register)
		r.POST("sendCode", apis.ApiSso.SendCode)
		r.POST("forgetPwd", apis.ApiSso.ForgetPwd)
	}

	rj := v1.Group("").Use(middleware.JwtHandler())
	{
		rj.POST("changePwd", apis.ApiSso.ChangePwd)
		rj.GET("myUserinfo", apis.ApiSso.MyUserInfo)
	}
}
