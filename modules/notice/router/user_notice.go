package router

import (
	"dilu/modules/notice/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerUserNoticeRouter)
}

// 默认需登录认证的路由
func registerUserNoticeRouter(v1 *gin.RouterGroup) {
	r := v1.Group("user-notice").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiUserNotice.Get)
		r.POST("/create", apis.ApiUserNotice.Create)
		r.POST("/update", apis.ApiUserNotice.Update)
		r.POST("/page", apis.ApiUserNotice.QueryPage)
		r.POST("/del", apis.ApiUserNotice.Del)
	}
}