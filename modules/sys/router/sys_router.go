package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckRoleRouter)
}

// 需要登录的路由
// func registerAuthCheckRoleRouter(v1 *gin.RouterGroup) {
// 	//api := apis.SSO{}
// 	r := v1.Group("/auth").Use(middleware.JWTAuthMiddleware())
// 	{
// 		r.POST("/myUserinfo", api.MyUserInfo)
// 		r.POST("/logout", api.Logout)
// 		r.POST("/changePwd", api.ChangePwd)
// 		r.POST("/bind", api.Bind)
// 		r.POST("/bindWechat", api.BindWechat)
// 		r.POST("/bindDing", api.BindDing)
// 		r.POST("/changeUserinfo", api.ChangeUserinfo)
// 	}
// }

// 无需认证的路由示例
func sysNoCheckRoleRouter(v1 *gin.RouterGroup) {
	api := apis.SysApi{}
	r := v1.Group("")
	{
		r.GET("/ping", api.Ping)
		r.GET("init", api.Init)
	}
}
